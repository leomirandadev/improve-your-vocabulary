package words

import (
	"context"

	"github.com/leomirandadev/improve-your-vocabulary/entities"
	"github.com/leomirandadev/improve-your-vocabulary/repositories"
	"github.com/leomirandadev/improve-your-vocabulary/utils/cache"
	"github.com/leomirandadev/improve-your-vocabulary/utils/logger"
)

type IService interface {
	Create(ctx context.Context, newWord entities.WordRequest) (*entities.Word, error)
	GetByID(ctx context.Context, ID uint64, userID uint64) (*entities.Word, error)
	GetAll(ctx context.Context, userID uint64) ([]entities.Word, error)
}

type services struct {
	repositories *repositories.Container
	log          logger.Logger
	cache        cache.Cache
}

func New(repo *repositories.Container, log logger.Logger, cache cache.Cache) IService {
	return &services{repositories: repo, log: log, cache: cache}
}

func (srv *services) Create(ctx context.Context, newWord entities.WordRequest) (*entities.Word, error) {
	id, err := srv.repositories.Word.Create(ctx, newWord)
	if err != nil {
		srv.log.ErrorContext(ctx, "Word.Service.Create", err)
		return nil, err
	}

	wordCreated, err := srv.repositories.Word.GetByID(ctx, id, newWord.UserID)
	if err != nil {
		srv.log.ErrorContext(ctx, "Word.Service.GetByID", err)
		return nil, err
	}

	return wordCreated, nil
}

func (srv *services) GetAll(ctx context.Context, userID uint64) ([]entities.Word, error) {

	var words []entities.Word
	if srv.cache.Get(ctx, CACHE_GET_ALL_WORDS, &words) {
		return words, nil
	}

	words, err := srv.repositories.Word.GetAll(ctx, userID)
	if err != nil {
		srv.log.ErrorContext(ctx, "Word.Service.GetAll", err)
		return nil, err
	}

	srv.cache.Set(ctx, CACHE_GET_ALL_WORDS, words)

	return words, nil
}

func (srv *services) GetByID(ctx context.Context, ID uint64, userID uint64) (*entities.Word, error) {

	wordWanted, err := srv.repositories.Word.GetByID(ctx, ID, userID)
	if err != nil {
		srv.log.ErrorContext(ctx, "Word.Service.GetByID", ID, err)
		return nil, err
	}

	return wordWanted, nil
}
