package services

import (
	"github.com/leomirandadev/improve-your-vocabulary/repositories"
	"github.com/leomirandadev/improve-your-vocabulary/services/user"
	"github.com/leomirandadev/improve-your-vocabulary/utils/cache"
	"github.com/leomirandadev/improve-your-vocabulary/utils/logger"
	"github.com/leomirandadev/improve-your-vocabulary/utils/mail"
)

// Container modelo para exportação dos serviços instanciados
type Container struct {
	User user.UserService
}

// Options struct de opções para a criação de uma instancia dos serviços
type Options struct {
	Repo  *repositories.Container
	Log   logger.Logger
	Cache cache.Cache
	Mail  mail.MailSender
}

// New cria uma nova instancia dos serviços
func New(opts Options) *Container {

	return &Container{
		User: user.New(opts.Repo, opts.Log),
	}
}
