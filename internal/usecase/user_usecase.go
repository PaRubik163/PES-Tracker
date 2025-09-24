package usecase

import (
	"errors"
	"tracker/internal/entity"
	"tracker/internal/repository"

	"gorm.io/gorm"
)

type UserUseCase struct{
	userRepo *repository.UserRepository
}

func NewUserUseCase(userRepo *repository.UserRepository) *UserUseCase {
	return &UserUseCase{
		userRepo: userRepo,
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

func (us *UserUseCase) Login(login, pass string) (*entity.User, error) {
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

	if err := us.userRepo.UpdateLogin(login); err != nil{
		return nil, err
	}
	
	return userDB, nil
}