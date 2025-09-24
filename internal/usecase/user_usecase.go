package usecase

import (
	"errors"
	"tracker/internal/dto"
	"tracker/internal/entity"
	"tracker/internal/repository"
	jwt "tracker/pkg/middleware"

	"gorm.io/gorm"
)

type UserUseCase struct{
	userRepo *repository.UserRepository
	redisRepo *repository.RedisRepo
	jwtService *jwt.Jwt
}

func NewUserUseCase(userRepo *repository.UserRepository, redis *repository.RedisRepo, jwt *jwt.Jwt) *UserUseCase {
	return &UserUseCase{
		userRepo: userRepo,
		redisRepo: redis,
		jwtService: jwt,
	}
}

func (us *UserUseCase) Register(login, pass string) error { 
	user := &entity.User{
		Login: login,
		Password: pass,
	}

	if err := user.CheckLoginAndPassword(user.Login, user.Password); err != nil{
		return err
	}

	if _, err := us.userRepo.GetByLogin(login); err == nil{
		return errors.New("user already exists")
	}
	
	if err := user.HashPassword(); err != nil{
		return err
	}

	err := us.userRepo.Create(user)
	
	if err != nil{
		return errors.New("failed to create new user")
	}
	
	return nil	
}

func (us *UserUseCase) Login(login, pass string) (*dto.UserSession, error) {
	userDB, err := us.userRepo.GetByLogin(login)

	if err != nil{
		if errors.Is(err, gorm.ErrRecordNotFound){
			return nil, errors.New("user doesn't exists")
		}
		return nil, err
	}

	if err := userDB.CheckLoginAndPassword(login, pass); err != nil{
		return nil, err
	}

	if ok := userDB.CheckPassword(pass); !ok{
		return nil, errors.New("invalid password")
	}

	resp, err := us.jwtService.GenerateToken()
	if err != nil{
		return nil, err
	}

	session := &dto.UserSession{
		ID: userDB.ID,
		Login: userDB.Login,
		Jwt: resp.Token,
		RegisteredAt: userDB.RegisteredAt,
		LastLogin: userDB.LastLogin,
	}

	err = us.redisRepo.SaveToken(resp.ID, session)
	if err != nil{
		return nil, err
	}

	if err := us.userRepo.UpdateLogin(login); err != nil{
		return nil, err
	}
	
	return session, nil
}