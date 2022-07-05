package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/leomirandadev/improve-your-vocabulary/repositories/meanings"
	"github.com/leomirandadev/improve-your-vocabulary/repositories/users"
	"github.com/leomirandadev/improve-your-vocabulary/repositories/words"
	"github.com/leomirandadev/improve-your-vocabulary/utils/logger"
)

// Container modelo para exportação dos repositórios instanciados
type Container struct {
	User    users.IRepository
	Word    words.IRepository
	Meaning meanings.IRepository
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
		User:    users.NewSqlx(opts.Log, opts.WriterSqlx, opts.ReaderSqlx),
		Word:    words.NewSqlx(opts.Log, opts.WriterSqlx, opts.ReaderSqlx),
		Meaning: meanings.NewSqlx(opts.Log, opts.WriterSqlx, opts.ReaderSqlx),
	}
}
