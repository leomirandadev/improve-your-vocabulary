package controllers

import (
	"github.com/leomirandadev/improve-your-vocabulary/internal/controllers/meanings"
	"github.com/leomirandadev/improve-your-vocabulary/internal/controllers/users"
	"github.com/leomirandadev/improve-your-vocabulary/internal/controllers/words"
	"github.com/leomirandadev/improve-your-vocabulary/internal/services"
	"github.com/leomirandadev/improve-your-vocabulary/pkg/logger"
	"github.com/leomirandadev/improve-your-vocabulary/pkg/token"
)

// Container modelo para exportação dos serviços instanciados
type Container struct {
	User    users.IController
	Word    words.IController
	Meaning meanings.IController
}

// Options struct de opções para a criação de uma instancia dos serviços
type Options struct {
	Srv   *services.Container
	Log   logger.Logger
	Token token.TokenHash
}

// New cria uma nova instancia dos serviços
func New(opts Options) *Container {
	return &Container{
		User:    users.New(opts.Srv, opts.Log, opts.Token),
		Word:    words.New(opts.Srv, opts.Log, opts.Token),
		Meaning: meanings.New(opts.Srv, opts.Log, opts.Token),
	}
}
