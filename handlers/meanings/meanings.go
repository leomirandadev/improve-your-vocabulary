package meanings

import (
	"github.com/leomirandadev/improve-your-vocabulary/controllers"
	"github.com/leomirandadev/improve-your-vocabulary/handlers/middlewares"
	"github.com/leomirandadev/improve-your-vocabulary/utils/httpRouter"
)

func New(mid *middlewares.Middleware, router httpRouter.Router, Ctrl *controllers.Container) {

	router.POST("/meanings", mid.Auth.Private(Ctrl.Meaning.Create))
	router.GET("/meanings", mid.Auth.Private(Ctrl.Meaning.GetAll))
	router.GET("/meanings/{id}", mid.Auth.Private(Ctrl.Meaning.GetByID))
}
