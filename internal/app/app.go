package app

import (
	"fmt"
	"tracker/internal/config"
	"tracker/internal/delivery/http/handler"
	"tracker/internal/delivery/http/route"
	"tracker/internal/repository"
	"tracker/internal/usecase"
	jwt "tracker/pkg/jwt"

	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type App struct{
	Conf *config.Config
	DB *gorm.DB
	Redis *repository.RedisRepo
	Jwt *jwt.Jwt
	Page *handler.PageHandler
	UserRepo *repository.UserRepository
	UserUseCase *usecase.UserUseCase
	UserHandler *handler.UserHandler
	SubscriptionRepo *repository.SubscriptionRepo
	SubscriptionUseCase *usecase.SubscriptionUseCase
	SubscriptionHandler *handler.SubscriptionHandler
	IncomeRepo *repository.IncomeRepository
	IncomeUseCase *usecase.IncomeUseCase
	IncomeHandler *handler.IncomeHandler
	ExpenseRepo *repository.ExpenseRepository
	ExpenseUseCase *usecase.ExpenseUseCase
	ExpenseHandler *handler.ExpenseHandler
	Router *route.Router
}

func NewApp(c *config.Config) (*App, error) {

	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", 
						c.DataBaseHost, 
						c.DataBasePort, 
						c.DataBaseUser, 
						c.DataBasePass,
						c.DataBaseName)
	
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil{
		return nil, err
	}
	redisRepo := repository.NewRedisRepo(c) //redis
	jwtService := jwt.NewJwt([]byte(c.JWTKey)) //jwt

	pgh := handler.NewPageHandler() //here processing pages

	//user section
	userRepo, err := repository.NewUserRepository(db)

	if err != nil{
		return nil, err
	}

	userUseCase := usecase.NewUserUseCase(userRepo, redisRepo, jwtService)
	userHandler := handler.NewUserHandler(userUseCase)

	//subscription section
	subRepo, err := repository.NewSubscriptionRepo(db)

	if err != nil{
		return nil, err
	}
	
	subUseCase := usecase.NewSubscriptionUseCase(subRepo)
	subHandler := handler.NewSubscriptionHandler(subUseCase)

	//income section
	inRepo, err := repository.NewIncomeRepository(db)

	if err != nil{
		return nil, err
	}

	inUseCase := usecase.NewIncomeUseCase(inRepo)
	inHandler := handler.NewIncomeHandler(inUseCase)

	//expense section
	expRepo, err := repository.NewExpenseRepository(db)
	
	if err != nil{
		return nil, err
	}

	expUseCase := usecase.NewExpenseUseCase(expRepo)
	expHandler := handler.NewExpenseHandler(expUseCase)
	//gin section
	router := route.NewRouter(pgh, userHandler, subHandler, inHandler, expHandler, jwtService) 
	router.SetupRouter()
	
	a := &App{
		Conf: c,
		DB: db,
		Redis: redisRepo,
		Jwt: jwtService,
		Page: pgh,
		UserRepo: userRepo,
		UserUseCase: userUseCase,
		UserHandler: userHandler,
		SubscriptionRepo: subRepo,
		SubscriptionUseCase: subUseCase,
		SubscriptionHandler: subHandler,
		IncomeRepo: inRepo,
		IncomeUseCase: inUseCase,
		IncomeHandler: inHandler,
		ExpenseRepo: expRepo,
		ExpenseUseCase: expUseCase,
		ExpenseHandler: expHandler,
		Router: router,
	}
	return a, nil
}

func (a *App) Run() error {
	logrus.Info("Start new app")
	return a.Router.Engine.Run(a.Conf.GinAddr)
}
