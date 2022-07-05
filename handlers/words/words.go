package words

import (
	"github.com/leomirandadev/improve-your-vocabulary/controllers"
	"github.com/leomirandadev/improve-your-vocabulary/handlers/middlewares"
	"github.com/leomirandadev/improve-your-vocabulary/utils/httpRouter"
)

func New(mid *middlewares.Middleware, router httpRouter.Router, Ctrl *controllers.Container) {

	router.POST("/words", mid.Auth.Public(Ctrl.Word.Create))
	router.GET("/words", mid.Auth.Public(Ctrl.Word.GetAll))
	router.GET("/words/{id}", mid.Auth.Public(Ctrl.Word.GetByID))
}
