package main

import (
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/leomirandadev/improve-your-vocabulary/configs"
	"github.com/leomirandadev/improve-your-vocabulary/controllers"
	"github.com/leomirandadev/improve-your-vocabulary/handlers"
	"github.com/leomirandadev/improve-your-vocabulary/repositories"
	"github.com/leomirandadev/improve-your-vocabulary/services"
	"github.com/leomirandadev/improve-your-vocabulary/utils/cache"
	"github.com/leomirandadev/improve-your-vocabulary/utils/httpRouter"
	"github.com/leomirandadev/improve-your-vocabulary/utils/logger"
	"github.com/leomirandadev/improve-your-vocabulary/utils/token"
	"github.com/leomirandadev/improve-your-vocabulary/utils/tracer"
	"github.com/leomirandadev/improve-your-vocabulary/utils/tracer/otel_jaeger"
)

// @securityDefinitions.apikey ApiKeyAuth
// @in header
// @name Authorization
func main() {

	router, log, tokenGenerator, cacheStore, tr, configs := toolsInit()
	defer tr.Close()

	repo := repositories.New(repositories.Options{
		Log:        log,
		ReaderSqlx: sqlx.MustConnect("mysql", configs.Database.Reader),
		WriterSqlx: sqlx.MustConnect("mysql", configs.Database.Writer),
		Cache:      cacheStore,
	})

	srv := services.New(services.Options{
		Log:  log,
		Repo: repo,
	})

	ctrl := controllers.New(controllers.Options{
		Log:   log,
		Srv:   srv,
		Token: tokenGenerator,
	})

	handlers.New(handlers.Options{
		Ctrl:   ctrl,
		Router: router,
		Log:    log,
		Token:  tokenGenerator,
	})

	router.SERVE(configs.Port)
}

func toolsInit() (httpRouter.Router, logger.Logger, token.TokenHash, cache.Cache, tracer.ITracer, configs.Config) {
	configs, err := configs.LoadConfig(".")
	if err != nil {
		log.Fatal("configs not loaded", err)
	}

	router := httpRouter.NewChiRouter()

	log := logger.NewLogrusLog()

	tokenGenerator := token.NewJWT()

	cacheStore := cache.NewMemcache(configs.Cache, log)

	tr := tracer.New(
		otel_jaeger.NewCollector(configs.Tracer),
	)

	return router, log, tokenGenerator, cacheStore, tr, configs
}
