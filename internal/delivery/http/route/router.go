package route

import (
	"tracker/internal/delivery/http/handler"
	"tracker/pkg/jwt"
	"tracker/internal/delivery/middleware"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Router struct{
	Engine *gin.Engine
	pageHandler *handler.PageHandler
	userHandler *handler.UserHandler
	subscriptionHandler *handler.SubscriptionHandler
	incomeHandler *handler.IncomeHandler
	expenseHandler *handler.ExpenseHandler
	jwtService *jwt.Jwt
}

func NewRouter(ph *handler.PageHandler, uh *handler.UserHandler, sh *handler.SubscriptionHandler,inH *handler.IncomeHandler, expH *handler.ExpenseHandler,jwt *jwt.Jwt) *Router {
	router := gin.Default()
	router.LoadHTMLGlob("./frontend/templates/*")
	router.Static("/static", "./frontend/static")
	return &Router{
		Engine: router,
		pageHandler: ph,
		userHandler: uh,
		subscriptionHandler: sh,
		incomeHandler: inH,
		expenseHandler: expH,
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
	r.Engine.GET("/new_income", r.pageHandler.NewIncome)
	r.Engine.GET("/income", r.pageHandler.GetAllIncome)
	r.Engine.GET("/new_expense", r.pageHandler.NewExpense)
	r.Engine.GET("/expenses", r.pageHandler.GetAllExpenses)
}

func (r *Router) setupAPIRoutes() {
	r.Engine.GET("/", func(c *gin.Context) {
    	c.Redirect(http.StatusFound, "/register")
	})

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

			auth.POST("/new_income",r.incomeHandler.HandlerAddIncome)
			auth.GET("/income", r.incomeHandler.HandlerGetAll)
			auth.DELETE("/income/:id", r.incomeHandler.DeleteIncome)

			auth.POST("/new_expense", r.expenseHandler.HandlerAdd)
			auth.GET("/expenses", r.expenseHandler.HandlerGetAll)
			auth.DELETE("/expense/:id", r.expenseHandler.HandlerDeleteExpense)
        }
    }
}