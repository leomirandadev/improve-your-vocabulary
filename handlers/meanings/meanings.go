package meanings

import (
	"github.com/leomirandadev/improve-your-vocabulary/controllers"
	"github.com/leomirandadev/improve-your-vocabulary/handlers/middlewares"
	"github.com/leomirandadev/improve-your-vocabulary/utils/httpRouter"
)

func New(mid *middlewares.Middleware, router httpRouter.Router, Ctrl *controllers.Container) {

	router.POST("/meanings", mid.Auth.Public(Ctrl.Meaning.Create))
	router.GET("/meanings", mid.Auth.Public(Ctrl.Meaning.GetAll))
	router.GET("/meanings/{id}", mid.Auth.Public(Ctrl.Meaning.GetByID))
}
