package usecase

import (
	"fmt"
	"errors"
	"time"
	"tracker/internal/dto"
	"tracker/internal/entity"
	"tracker/internal/logger"
	"tracker/internal/repository"
	jwt "tracker/pkg/jwt"
	pb "tracker/pkg/logger"
	"gorm.io/gorm"
)

type UserUseCase struct{
	userRepo *repository.UserRepository
	redisRepo *repository.RedisRepo
	jwtService *jwt.Jwt
	loggerAdapter *logger.Adapter
}

func NewUserUseCase(
	userRepo *repository.UserRepository, 
	redis *repository.RedisRepo, 
	jwt *jwt.Jwt, 
	loggerAdapter *logger.Adapter,
	) *UserUseCase {
	return &UserUseCase{
		userRepo: userRepo,
		redisRepo: redis,
		jwtService: jwt,
		loggerAdapter: loggerAdapter,
	}
}

func (us *UserUseCase) Register(login, pass string) error { 
	start := time.Now()

	user := &entity.User{
		Login: login,
		Password: pass,
		RegisteredAt: time.Now(),
		LastLogin: time.Now(),
		Role: "user",
	}

	if err := user.CheckLogin(user.Login); err != nil{
		return err
	}

	if err := user.CheckPassword(user.Password); err != nil{
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
		us.loggerAdapter.Users("ERROR", &pb.LogData{
			UserId: user.Login,
			Token: "",
			Method: "POST",
			Path: "/register",
			Status: "500 INTERNAL",
			Time: time.Now().Format(time.RFC3339),
			Latency: fmt.Sprintf("%v", time.Since(start)),
			Error: err.Error(),
		})

		return errors.New("failed to create new user")
	}
	
	us.loggerAdapter.Users("INFO", &pb.LogData{
		UserId: user.Login,
		Token: "",
		Method: "POST",
		Path: "/register",
		Status: "201 CREATED",
		Time: time.Now().Format(time.RFC3339),
		Latency: fmt.Sprintf("%v", time.Since(start)),
		Error: "",
	})
	return nil	
}

func (us *UserUseCase) Login(login, pass string) (*dto.UserSession, error) {
	start := time.Now()

	userDB, err := us.userRepo.GetByLogin(login)

	if err != nil{
		if errors.Is(err, gorm.ErrRecordNotFound){
			return nil, errors.New("user doesn't exists")
		}
		return nil, err
	}

	if err := userDB.CheckLogin(login); err != nil{
		return nil, err
	}

	if ok := userDB.CheckHashedPassword(pass); !ok{
		return nil, errors.New("invalid login or password")
	}

	resp, err := us.jwtService.GenerateToken(userDB.ID, userDB.Role)
	if err != nil{
		return nil, err
	}

	if err := us.userRepo.UpdateLogin(login); err != nil{
		return nil, err
	}
	
	session := &dto.UserSession{
		ID: userDB.ID,
		Login: userDB.Login,
		Token: resp.Token,
		CreateSessionAt: time.Now(),
		Role: resp.UserRole,
	}

	err = us.redisRepo.SaveUser(resp.ID, session)

	if err != nil{
		us.loggerAdapter.Users("ERROR", &pb.LogData{
			UserId: fmt.Sprintf("%d", session.ID),
			Token: resp.Token,
			Method: "POST",
			Path: "/login",
			Status: "500 INTERNAL",
			Time: start.Format(time.RFC3339),
			Latency: fmt.Sprintf("%v", time.Since(start)),
			Error: err.Error(),
		})
		return nil, err
	}
	
	us.loggerAdapter.Users("INFO", &pb.LogData{
		UserId: fmt.Sprintf("%d", session.ID),
		Token: resp.Token,
		Method: "POST",
		Path: "/login",
		Status: "200 OK",
		Time: start.Format(time.RFC3339),
		Latency: fmt.Sprintf("%v", time.Since(start)),
		Error: "",
	})
	return session, nil
}

func (us *UserUseCase) Logout(uuid string) error {
	start := time.Now()
	err := us.redisRepo.DeleteUser(uuid)

	if err != nil{
		us.loggerAdapter.Users("ERROR", &pb.LogData{
			UserId: uuid,
			Token: "",
			Method: "POST",
			Path: "/logout",
			Status: "500 INTERNAL",
			Time: start.Format(time.RFC3339),
			Latency: fmt.Sprintf("%v", time.Since(start)),
			Error: err.Error(),
		})
		return err
	}

	us.loggerAdapter.Users("INFO", &pb.LogData{
		UserId: uuid,
		Token: "",
		Method: "POST",
		Path: "/logout",
		Status: "200 OK",
		Time: start.Format(time.RFC3339),
		Latency: fmt.Sprintf("%v", time.Since(start)),
		Error: "",
	})

	return nil
}

func (us *UserUseCase) GetMe(uuid string) (*dto.UserSession, error) {
	start := time.Now()

	userSession, err := us.redisRepo.GetUser(uuid)

	if err != nil{
		us.loggerAdapter.Users("ERROR", &pb.LogData{
			UserId: fmt.Sprintf("%v", userSession.ID),
			Token: userSession.Token,
			Method: "GET",
			Path: "/me",
			Status: "500 INTERNAL",
			Time: start.Format(time.RFC3339),
			Latency: fmt.Sprintf("%v", time.Since(start)),
			Error: err.Error(),
		})
		return nil, err
	}

	subQuntity, err := us.userRepo.CountUsersSubscription(userSession.ID)
	if err != nil{
		us.loggerAdapter.Users("ERROR", &pb.LogData{
			UserId: fmt.Sprintf("%v", userSession.ID),
			Token: userSession.Token,
			Method: "GET",
			Path: "/me",
			Status: "500 INTERNAL",
			Time: start.Format(time.RFC3339),
			Latency: fmt.Sprintf("%v", time.Since(start)),
			Error: err.Error(),
		})
		return nil, err
	}

	incomeQuantity, err := us.userRepo.CountUsersIncome(userSession.ID)
	if err != nil{
		us.loggerAdapter.Users("ERROR", &pb.LogData{
			UserId: fmt.Sprintf("%v", userSession.ID),
			Token: userSession.Token,
			Method: "GET",
			Path: "/me",
			Status: "500 INTERNAL",
			Time: start.Format(time.RFC3339),
			Latency: fmt.Sprintf("%v", time.Since(start)),
			Error: err.Error(),
		})
		return nil, err
	}

	expensesQuantity, err := us.userRepo.CountUserExpenses(userSession.ID)
	if err != nil{
		us.loggerAdapter.Users("ERROR", &pb.LogData{
			UserId: fmt.Sprintf("%v", userSession.ID),
			Token: userSession.Token,
			Method: "GET",
			Path: "/me",
			Status: "500 INTERNAL",
			Time: start.Format(time.RFC3339),
			Latency: fmt.Sprintf("%v", time.Since(start)),
			Error: err.Error(),
		})
		return nil, err
	}

	userSession.SubscriptionsQuantity = subQuntity
	userSession.IncomeMonth = incomeQuantity
	userSession.ExpensesMonth = expensesQuantity

	us.loggerAdapter.Users("INFO", &pb.LogData{
		UserId: fmt.Sprintf("%v", userSession.ID),
		Token: userSession.Token,
		Method: "GET",
		Path: "/me",
		Status: "200 OK",
		Time: start.Format(time.RFC3339),
		Latency: fmt.Sprintf("%v", time.Since(start)),
		Error: "",
	})

	return userSession, nil
}

func (us *UserUseCase) GettAllUsers() ([]entity.User, error){
	return us.userRepo.GetAllUsers()
}

func (us *UserUseCase) GetUserByID(userID int) (*entity.User, error){
	return us.userRepo.GetUserByID(userID)
}