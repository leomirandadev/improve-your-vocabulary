package handlers

import (
	"github.com/leomirandadev/improve-your-vocabulary/internal/controllers"
	"github.com/leomirandadev/improve-your-vocabulary/internal/handlers/meanings"
	"github.com/leomirandadev/improve-your-vocabulary/internal/handlers/middlewares"
	"github.com/leomirandadev/improve-your-vocabulary/internal/handlers/swagger"
	"github.com/leomirandadev/improve-your-vocabulary/internal/handlers/users"
	"github.com/leomirandadev/improve-your-vocabulary/internal/handlers/words"
	"github.com/leomirandadev/improve-your-vocabulary/pkg/httpRouter"
	"github.com/leomirandadev/improve-your-vocabulary/pkg/logger"
	"github.com/leomirandadev/improve-your-vocabulary/pkg/token"
)

type Options struct {
	Ctrl   *controllers.Container
	Log    logger.Logger
	Token  token.TokenHash
	Router httpRouter.Router
}

func New(opts Options) {

	mid := middlewares.New(opts.Token, opts.Log)

	users.New(mid, opts.Router, opts.Ctrl)
	words.New(mid, opts.Router, opts.Ctrl)
	meanings.New(mid, opts.Router, opts.Ctrl)
	swagger.New(opts.Router)
}
