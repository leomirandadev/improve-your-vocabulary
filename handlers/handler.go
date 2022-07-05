package handlers

import (
	"github.com/leomirandadev/improve-your-vocabulary/controllers"
	"github.com/leomirandadev/improve-your-vocabulary/handlers/middlewares"
	"github.com/leomirandadev/improve-your-vocabulary/handlers/users"
	"github.com/leomirandadev/improve-your-vocabulary/utils/httpRouter"
	"github.com/leomirandadev/improve-your-vocabulary/utils/logger"
	"github.com/leomirandadev/improve-your-vocabulary/utils/token"
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
}
