package middlewares

import (
	"github.com/leomirandadev/improve-your-vocabulary/handlers/middlewares/auth"
	"github.com/leomirandadev/improve-your-vocabulary/utils/logger"
	"github.com/leomirandadev/improve-your-vocabulary/utils/token"
)

type Middleware struct {
	Auth auth.AuthMiddleware
}

func New(tokenHasher token.TokenHash, log logger.Logger) *Middleware {
	return &Middleware{
		Auth: auth.NewBearer(tokenHasher, log),
	}
}
