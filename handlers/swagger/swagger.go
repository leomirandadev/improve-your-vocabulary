package swagger

import (
	docs "github.com/leomirandadev/improve-your-vocabulary/docs"
	"github.com/leomirandadev/improve-your-vocabulary/utils/httpRouter"
)

func New(router httpRouter.Router) {

	docs.SwaggerInfo.Title = "Swagger"
	docs.SwaggerInfo.Version = "1.0"
	docs.SwaggerInfo.Schemes = []string{"http", "https"}

	// router.GET("/swagger/*", httpSwagger.Handler(
	// 	httpSwagger.URL("doc.json"),
	// ))
}
