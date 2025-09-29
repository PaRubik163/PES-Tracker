package route

import (
	"tracker/internal/delivery/http/handler"
	"tracker/pkg/jwt"
	"tracker/internal/delivery/middleware"
	"github.com/gin-gonic/gin"
)

type Router struct{
	Engine *gin.Engine
	pageHandler *handler.PageHandler
	userHandler *handler.UserHandler
	subscriptionHandler *handler.SubscriptionHandler
	jwtService *jwt.Jwt
}

func NewRouter(ph *handler.PageHandler, uh *handler.UserHandler, sh *handler.SubscriptionHandler, jwt *jwt.Jwt) *Router {
	router := gin.Default()
	router.LoadHTMLGlob("./frontend/templates/*")
	return &Router{
		Engine: router,
		pageHandler: ph,
		userHandler: uh,
		subscriptionHandler: sh,
		jwtService: jwt,
	}
}

func (r *Router) SetupRouter() {
    r.setupPageRoutes()
    r.setupAPIRoutes()
}

func (r *Router) setupPageRoutes() {
    r.Engine.GET("/login", r.pageHandler.Login)
    r.Engine.GET("/register", r.pageHandler.Register)
	r.Engine.GET("/me", r.pageHandler.GetMe)
	r.Engine.GET("/new_subscription", r.pageHandler.NewSubscription)
    r.Engine.GET("/subscriptions", r.pageHandler.GetAllSubscriptions)
}

func (r *Router) setupAPIRoutes() {
    api := r.Engine.Group("/api/v1")
    {
        api.POST("/register", r.userHandler.HandlerRegister)
        api.POST("/login", r.userHandler.HandlerLogin)

        auth := api.Group("/")
        auth.Use(middleware.AuthMiddleware(r.jwtService))
        {
            auth.POST("/logout", r.userHandler.HandlerLogout)
            auth.GET("/me", r.userHandler.HandlerGetMe)

            auth.POST("/new_subscription", r.subscriptionHandler.HandlerAdd)
            auth.GET("/subscriptions", r.subscriptionHandler.HandlerGetAll)
			auth.DELETE("/subscription/:id", r.subscriptionHandler.HandlerDeleteSubscription)
        }
    }
}