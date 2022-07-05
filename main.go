package main

import (
	"os"
	"strconv"
	"time"

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

	router, log, tokenHasher, cacheStore, _ := toolsInit()

	repo := repositories.New(repositories.Options{
		Log:        log,
		ReaderSqlx: configs.GetReaderSqlx(),
		WriterSqlx: configs.GetWriterSqlx(),
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

func toolsInit() (httpRouter.Router, logger.Logger, token.TokenHash, cache.Cache, mail.MailSender) {

	router := httpRouter.NewMuxRouter()
	log := logger.NewLogrusLog()
	tokenHasher := token.NewJWT()

	cacheStore := cache.NewMemcache(cache.Options{
		URL: os.Getenv("CACHE_URL"),
		Expiration: func() time.Duration {
			cacheExpiration, _ := strconv.ParseInt(os.Getenv("CACHE_EXP"), 10, 64)
			return time.Duration(cacheExpiration)
		}(),
	}, log)

	return router, log, tokenHasher, cacheStore, nil
}
