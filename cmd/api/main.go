package main

import (
	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx"
	"github.com/leomirandadev/improve-your-vocabulary/internal/controllers"
	"github.com/leomirandadev/improve-your-vocabulary/internal/handlers"
	"github.com/leomirandadev/improve-your-vocabulary/internal/repositories"
	"github.com/leomirandadev/improve-your-vocabulary/internal/services"
	"github.com/leomirandadev/improve-your-vocabulary/pkg/cache"
	"github.com/leomirandadev/improve-your-vocabulary/pkg/envs"
	"github.com/leomirandadev/improve-your-vocabulary/pkg/httpRouter"
	"github.com/leomirandadev/improve-your-vocabulary/pkg/logger"
	"github.com/leomirandadev/improve-your-vocabulary/pkg/token"
	"github.com/leomirandadev/improve-your-vocabulary/pkg/tracer"
	"github.com/leomirandadev/improve-your-vocabulary/pkg/tracer/otel_jaeger"
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

type Config struct {
	Port     string              `mapstructure:"port"`
	Env      string              `mapstructure:"env"`
	Cache    cache.Options       `mapstructure:"cache"`
	Tracer   otel_jaeger.Options `mapstructure:"tracer"`
	Database struct {
		Reader string `mapstructure:"reader"`
		Writer string `mapstructure:"writer"`
	} `mapstructure:"database"`
}

func toolsInit() (httpRouter.Router, logger.Logger, token.TokenHash, cache.Cache, tracer.ITracer, Config) {
	log := logger.NewLogrusLog()

	var configs Config
	err := envs.Load(".", &configs)
	if err != nil {
		log.Fatal("configs not loaded", err)
	}

	router := httpRouter.NewChiRouter()

	tokenGenerator := token.NewJWT()

	cacheStore := cache.NewMemcache(configs.Cache, log)

	tr := tracer.New(
		otel_jaeger.NewCollector(configs.Tracer),
	)

	return router, log, tokenGenerator, cacheStore, tr, configs
}
