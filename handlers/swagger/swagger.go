package swagger

import (
	docs "github.com/leomirandadev/improve-your-vocabulary/docs"
	"github.com/leomirandadev/improve-your-vocabulary/utils/httpRouter"
	httpSwagger "github.com/swaggo/http-swagger"
)

func New(router httpRouter.Router) {

	docs.SwaggerInfo.Title = "Swagger"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	router.GET("/swagger/*", router.ParseHandler(
		httpSwagger.Handler(
			httpSwagger.URL("doc.json"),
		),
	))
}
