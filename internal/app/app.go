package app

import (
	"fmt"
	"tracker/internal/config"
	"tracker/internal/handler"
	jwt "tracker/pkg/middleware"
	"tracker/internal/repository"
	"tracker/internal/usecase"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type App struct{
	Conf *config.Config
	DB *gorm.DB
	UserRepo *repository.UserRepository
	UserUseCase *usecase.UserUseCase
	UserHandler *handler.UserHandler
	Router *gin.Engine
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

	ur := repository.NewUserRepository(db)
	redisRepo := repository.NewRedisRepo(c)
	jwtService := jwt.NewJwt([]byte(c.JWTKey))
	
	us := usecase.NewUserUseCase(ur,redisRepo, jwtService)
	uh := handler.NewUserHandler(us)

	router := gin.Default()
	

	a := &App{
		Conf: c,
		DB: db,
		UserRepo: ur,
		UserUseCase: us,
		UserHandler: uh,
		Router: router,
	}

	a.setupRouter()

	return a, nil
}

func (a *App) Run() error {
	logrus.Info("Start new app")
	return a.Router.Run(a.Conf.GinAddr)
}

func (a *App) setupRouter(){
	api := a.Router.Group("/api/v1")
	api.POST("/register", a.UserHandler.HandlerRegister) //user registration
	api.POST("/login", a.UserHandler.HandlerLogin)	//user login
}