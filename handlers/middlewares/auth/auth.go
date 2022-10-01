package auth

import (
	"github.com/leomirandadev/improve-your-vocabulary/utils/httpRouter"
)

type AuthMiddleware interface {
	Public(next httpRouter.HandlerFunc) httpRouter.HandlerFunc
	Private(next httpRouter.HandlerFunc) httpRouter.HandlerFunc
	Admin(next httpRouter.HandlerFunc) httpRouter.HandlerFunc
}
