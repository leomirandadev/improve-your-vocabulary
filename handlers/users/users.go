package users

import (
	"github.com/leomirandadev/improve-your-vocabulary/controllers"
	"github.com/leomirandadev/improve-your-vocabulary/handlers/middlewares"
	"github.com/leomirandadev/improve-your-vocabulary/utils/httpRouter"
)

func New(mid *middlewares.Middleware, router httpRouter.Router, Ctrl *controllers.Container) {

	router.POST("/users", mid.Auth.Admin(Ctrl.User.Create))
	router.GET("/users/{id}", mid.Auth.Admin(Ctrl.User.GetByID))
	router.POST("/users/auth", mid.Auth.Public(Ctrl.User.Auth))

}
