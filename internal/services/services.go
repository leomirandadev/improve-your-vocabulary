package services

import (
	"github.com/leomirandadev/improve-your-vocabulary/internal/repositories"
	"github.com/leomirandadev/improve-your-vocabulary/internal/services/meanings"
	"github.com/leomirandadev/improve-your-vocabulary/internal/services/users"
	"github.com/leomirandadev/improve-your-vocabulary/internal/services/words"
	"github.com/leomirandadev/improve-your-vocabulary/pkg/logger"
)

// Container modelo para exportação dos serviços instanciados
type Container struct {
	User    users.IService
	Word    words.IService
	Meaning meanings.IService
}

// Options struct de opções para a criação de uma instancia dos serviços
type Options struct {
	Repo *repositories.Container
	Log  logger.Logger
}

// New cria uma nova instancia dos serviços
func New(opts Options) *Container {
	return &Container{
		User:    users.New(opts.Repo, opts.Log),
		Word:    words.New(opts.Repo, opts.Log),
		Meaning: meanings.New(opts.Repo, opts.Log),
	}
}
