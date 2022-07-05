package controllers

import (
	"github.com/leomirandadev/improve-your-vocabulary/controllers/meaning"
	"github.com/leomirandadev/improve-your-vocabulary/controllers/user"
	"github.com/leomirandadev/improve-your-vocabulary/controllers/word"
	"github.com/leomirandadev/improve-your-vocabulary/services"
	"github.com/leomirandadev/improve-your-vocabulary/utils/logger"
	"github.com/leomirandadev/improve-your-vocabulary/utils/token"
)

// Container modelo para exportação dos serviços instanciados
type Container struct {
	User    user.IController
	Word    word.IController
	Meaning meaning.IController
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
		User:    user.New(opts.Srv, opts.Log, opts.Token),
		Word:    word.New(opts.Srv, opts.Log, opts.Token),
		Meaning: meaning.New(opts.Srv, opts.Log, opts.Token),
	}
}
