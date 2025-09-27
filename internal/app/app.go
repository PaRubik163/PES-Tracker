package app

import (
	"fmt"
	"tracker/internal/config"
	"tracker/internal/delivery/handler"
	"tracker/internal/delivery/middleware"
	"tracker/internal/repository"
	"tracker/internal/usecase"
	jwt "tracker/pkg/jwt"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type App struct{
	Conf *config.Config
	DB *gorm.DB
	Page *handler.PageHandler
	UserRepo *repository.UserRepository
	UserUseCase *usecase.UserUseCase
	UserHandler *handler.UserHandler
	SubscriptionRepo *repository.SubscriptionRepo
	SubscriptionUseCase *usecase.SubscriptionUseCase
	SubscriptionHandler *handler.SubscriptionHandler
	Router *gin.Engine
	Jwt *jwt.Jwt
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

	//gin section
	router := gin.Default()
	router.LoadHTMLGlob("frontend/templates/*")
	

	a := &App{
		Conf: c,
		DB: db,
		Page: pgh,
		UserRepo: ur,
		UserUseCase: us,
		UserHandler: uh,
		SubscriptionRepo: sr,
		SubscriptionUseCase: sUseCase,
		SubscriptionHandler: sh,
		Router: router,
		Jwt: jwtService,
	}

	a.setupRouter()

	return a, nil
}

func (a *App) Run() error {
	logrus.Info("Start new app")
	return a.Router.Run(a.Conf.GinAddr)
}

func (a *App) setupRouter(){
	a.Router.GET("/register", a.Page.Register)
	a.Router.GET("/login", a.Page.Login)
	a.Router.GET("/me", a.Page.GetMe)
	a.Router.GET("/subscriptions", a.Page.GetAllSubscriptions)

	api := a.Router.Group("/api/v1")
	api.POST("/register", a.UserHandler.HandlerRegister) //user registration
	api.POST("/login", a.UserHandler.HandlerLogin)	//user login	

	auth := api.Group("/")
	auth.Use(middleware.AuthMiddleware(a.Jwt))
	auth.POST("/logout", a.UserHandler.HandlerLogout) //user logout
	auth.GET("/me", a.UserHandler.HandlerGetMe) //user information

	auth.POST("/new_subscriptions", a.SubscriptionHandler.HandlerAdd) //add subscription
	auth.GET("/subscriptions", a.SubscriptionHandler.HandlerGetAll)
}