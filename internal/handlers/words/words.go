package words

import (
	"github.com/leomirandadev/improve-your-vocabulary/internal/controllers"
	"github.com/leomirandadev/improve-your-vocabulary/internal/handlers/middlewares"
	"github.com/leomirandadev/improve-your-vocabulary/pkg/httpRouter"
)

func New(mid *middlewares.Middleware, router httpRouter.Router, Ctrl *controllers.Container) {

	router.POST("/words", mid.Auth.Private(Ctrl.Word.Create))
	router.GET("/words", mid.Auth.Private(Ctrl.Word.GetAll))
	router.GET("/words/{id}", mid.Auth.Private(Ctrl.Word.GetByID))
}
