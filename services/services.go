package services

import (
	"github.com/leomirandadev/improve-your-vocabulary/repositories"
	"github.com/leomirandadev/improve-your-vocabulary/services/meaning"
	"github.com/leomirandadev/improve-your-vocabulary/services/user"
	"github.com/leomirandadev/improve-your-vocabulary/services/word"
	"github.com/leomirandadev/improve-your-vocabulary/utils/cache"
	"github.com/leomirandadev/improve-your-vocabulary/utils/logger"
)

// Container modelo para exportação dos serviços instanciados
type Container struct {
	User    user.IService
	Word    word.IService
	Meaning meaning.IService
}

// Options struct de opções para a criação de uma instancia dos serviços
type Options struct {
	Repo  *repositories.Container
	Log   logger.Logger
	Cache cache.Cache
}

// New cria uma nova instancia dos serviços
func New(opts Options) *Container {
	return &Container{
		User:    user.New(opts.Repo, opts.Log),
		Word:    word.New(opts.Repo, opts.Log, opts.Cache),
		Meaning: meaning.New(opts.Repo, opts.Log, opts.Cache),
	}
}
