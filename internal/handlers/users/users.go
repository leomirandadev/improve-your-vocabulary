package users

import (
	"github.com/leomirandadev/improve-your-vocabulary/internal/controllers"
	"github.com/leomirandadev/improve-your-vocabulary/internal/handlers/middlewares"
	"github.com/leomirandadev/improve-your-vocabulary/pkg/httpRouter"
)

func New(mid *middlewares.Middleware, router httpRouter.Router, Ctrl *controllers.Container) {

	router.POST("/users", mid.Auth.Public(Ctrl.User.Create))
	router.GET("/users/{id}", mid.Auth.Admin(Ctrl.User.GetByID))
	router.POST("/users/auth", mid.Auth.Public(Ctrl.User.Auth))
}
