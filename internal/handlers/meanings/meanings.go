package meanings

import (
	"github.com/leomirandadev/improve-your-vocabulary/internal/controllers"
	"github.com/leomirandadev/improve-your-vocabulary/internal/handlers/middlewares"
	"github.com/leomirandadev/improve-your-vocabulary/pkg/httpRouter"
)

func New(mid *middlewares.Middleware, router httpRouter.Router, Ctrl *controllers.Container) {

	router.POST("/meanings", mid.Auth.Private(Ctrl.Meaning.Create))
	router.GET("/meanings", mid.Auth.Private(Ctrl.Meaning.GetAll))
	router.GET("/meanings/{id}", mid.Auth.Private(Ctrl.Meaning.GetByID))
}
