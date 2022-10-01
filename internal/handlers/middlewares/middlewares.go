package middlewares

import (
	"github.com/leomirandadev/improve-your-vocabulary/internal/handlers/middlewares/auth"
	"github.com/leomirandadev/improve-your-vocabulary/pkg/logger"
	"github.com/leomirandadev/improve-your-vocabulary/pkg/token"
)

type Middleware struct {
	Auth auth.AuthMiddleware
}

func New(tokenHasher token.TokenHash, log logger.Logger) *Middleware {
	return &Middleware{
		Auth: auth.NewBearer(tokenHasher, log),
	}
}
