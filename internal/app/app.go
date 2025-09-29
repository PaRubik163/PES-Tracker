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
	ur := repository.NewUserRepository(db)
	us := usecase.NewUserUseCase(ur,redisRepo, jwtService)
	uh := handler.NewUserHandler(us)

	//subscription section
	sr := repository.NewSubscriptionRepo(db)
	sUseCase := usecase.NewSubscriptionUseCase(sr)
	sh := handler.NewSubscriptionHandler(sUseCase)

	router := route.NewRouter(pgh, uh, sh, jwtService) //gin section
	router.SetupRouter()
	
	a := &App{
		Conf: c,
		DB: db,
		Redis: redisRepo,
		Jwt: jwtService,
		Page: pgh,
		UserRepo: ur,
		UserUseCase: us,
		UserHandler: uh,
		SubscriptionRepo: sr,
		SubscriptionUseCase: sUseCase,
		SubscriptionHandler: sh,
		Router: router,
	}
	return a, nil
}

func (a *App) Run() error {
	logrus.Info("Start new app")
	return a.Router.Engine.Run(a.Conf.GinAddr)
}
