package swagger

import (
	"github.com/leomirandadev/improve-your-vocabulary/controllers"
	docs "github.com/leomirandadev/improve-your-vocabulary/docs"
	"github.com/leomirandadev/improve-your-vocabulary/handlers/middlewares"
	"github.com/leomirandadev/improve-your-vocabulary/utils/httpRouter"
	httpSwagger "github.com/swaggo/http-swagger"
)

func New(mid *middlewares.Middleware, router httpRouter.Router, Ctrl *controllers.Container) {

	docs.SwaggerInfo.Title = "Swagger"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	router.GET("/swagger/*", httpSwagger.Handler(
		httpSwagger.URL("doc.json"),
	))
}
