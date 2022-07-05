package repositories

import (
	"github.com/jmoiron/sqlx"
	"github.com/leomirandadev/improve-your-vocabulary/repositories/meaning"
	"github.com/leomirandadev/improve-your-vocabulary/repositories/user"
	"github.com/leomirandadev/improve-your-vocabulary/repositories/word"
	"github.com/leomirandadev/improve-your-vocabulary/utils/logger"
)

// Container modelo para exportação dos repositórios instanciados
type Container struct {
	User    user.IRepository
	Word    word.IRepository
	Meaning meaning.IRepository
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
		User:    user.NewSqlx(opts.Log, opts.WriterSqlx, opts.ReaderSqlx),
		Word:    word.NewSqlx(opts.Log, opts.WriterSqlx, opts.ReaderSqlx),
		Meaning: meaning.NewSqlx(opts.Log, opts.WriterSqlx, opts.ReaderSqlx),
	}
}
