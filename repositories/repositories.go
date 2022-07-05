package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/leomirandadev/improve-your-vocabulary/repositories/user"
	"github.com/leomirandadev/improve-your-vocabulary/utils/logger"
)

// Container modelo para exportação dos repositórios instanciados
type Container struct {
	User user.UserRepository
}

// Options struct de opções para a criação de uma instancia dos serviços
type Options struct {
	Log        logger.Logger
	WriterSqlx *sqlx.DB
	ReaderSqlx *sqlx.DB
}

// New cria uma nova instancia dos repositórios
func New(opts Options) *Container {
	return &Container{
		User: user.NewSqlxRepository(opts.Log, opts.WriterSqlx, opts.ReaderSqlx),
	}
}
