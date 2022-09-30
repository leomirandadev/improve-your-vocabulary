package repositories

import (
	"github.com/jmoiron/sqlx"
	meaningsCache "github.com/leomirandadev/improve-your-vocabulary/repositories/cache/meanings"
	wordsCache "github.com/leomirandadev/improve-your-vocabulary/repositories/cache/words"
	"github.com/leomirandadev/improve-your-vocabulary/repositories/database/meanings"
	"github.com/leomirandadev/improve-your-vocabulary/repositories/database/users"
	"github.com/leomirandadev/improve-your-vocabulary/repositories/database/words"
	"github.com/leomirandadev/improve-your-vocabulary/utils/cache"
	"github.com/leomirandadev/improve-your-vocabulary/utils/logger"
)

// Container modelo para exportação dos repositórios instanciados
type Container struct {
	Database SqlContainer
	Cache    CacheContainer
}

type SqlContainer struct {
	User    users.IRepository
	Word    words.IRepository
	Meaning meanings.IRepository
}

type CacheContainer struct {
	Word    wordsCache.ICache
	Meaning meaningsCache.ICache
}

// Options struct de opções para a criação de uma instancia dos serviços
type Options struct {
	Log        logger.Logger
	WriterSqlx *sqlx.DB
	ReaderSqlx *sqlx.DB
	Cache      cache.Cache
}

// New cria uma nova instancia dos repositórios
func New(opts Options) *Container {
	return &Container{
		Database: SqlContainer{
			User:    users.NewSqlx(opts.Log, opts.WriterSqlx, opts.ReaderSqlx),
			Word:    words.NewSqlx(opts.Log, opts.WriterSqlx, opts.ReaderSqlx),
			Meaning: meanings.NewSqlx(opts.Log, opts.WriterSqlx, opts.ReaderSqlx),
		},
		Cache: CacheContainer{
			Word:    wordsCache.NewCache(opts.Cache),
			Meaning: meaningsCache.NewCache(opts.Cache),
		},
	}
}
