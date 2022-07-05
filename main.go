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
	"github.com/leomirandadev/improve-your-vocabulary/utils/mail"
	"github.com/leomirandadev/improve-your-vocabulary/utils/token"
)

func main() {

	configs, err := configs.LoadConfig(".")
	if err != nil {
		log.Fatal("configs not loaded")
	}

	router, log, tokenHasher, cacheStore, _ := toolsInit(configs.Cache)

	repo := repositories.New(repositories.Options{
		Log:        log,
		ReaderSqlx: sqlx.MustConnect("mysql", configs.Database.Reader),
		WriterSqlx: sqlx.MustConnect("mysql", configs.Database.Writer),
	})

	srv := services.New(services.Options{
		Log:   log,
		Repo:  repo,
		Cache: cacheStore,
	})

	ctrl := controllers.New(controllers.Options{
		Log:   log,
		Srv:   srv,
		Token: tokenHasher,
	})

	handlers.New(handlers.Options{
		Ctrl:   ctrl,
		Router: router,
		Log:    log,
		Token:  tokenHasher,
	})

	router.SERVE(":8080")
}

func toolsInit(cacheConfig configs.ConfigCache) (httpRouter.Router, logger.Logger, token.TokenHash, cache.Cache, mail.MailSender) {

	router := httpRouter.NewMuxRouter()
	log := logger.NewLogrusLog()
	tokenHasher := token.NewJWT()

	cacheStore := cache.NewMemcache(cache.Options{
		URL:        cacheConfig.URL,
		Expiration: cacheConfig.Expiration,
	}, log)

	return router, log, tokenHasher, cacheStore, nil
}
