package meaning

import (
	"context"

	"github.com/leomirandadev/improve-your-vocabulary/entities"
	"github.com/leomirandadev/improve-your-vocabulary/repositories"
	"github.com/leomirandadev/improve-your-vocabulary/utils/cache"
	"github.com/leomirandadev/improve-your-vocabulary/utils/logger"
)

type IService interface {
	Create(ctx context.Context, newMeaning entities.MeaningRequest) (*entities.Meaning, error)
	GetByID(ctx context.Context, ID uint64) (*entities.Meaning, error)
	GetAll(ctx context.Context) ([]entities.Meaning, error)
}

type services struct {
	repositories *repositories.Container
	log          logger.Logger
	cache        cache.Cache
}

func New(repo *repositories.Container, log logger.Logger, cache cache.Cache) IService {
	return &services{repositories: repo, log: log, cache: cache}
}

func (srv *services) Create(ctx context.Context, newMeaning entities.MeaningRequest) (*entities.Meaning, error) {
	id, err := srv.repositories.Meaning.Create(ctx, newMeaning)
	if err != nil {
		srv.log.ErrorContext(ctx, "Meaning.Service.Create", err)
		return nil, err
	}

	meaningCreated, err := srv.repositories.Meaning.GetByID(ctx, id)
	if err != nil {
		srv.log.ErrorContext(ctx, "Meaning.Service.GetByID", err)
		return nil, err
	}

	return meaningCreated, nil
}

func (srv *services) GetAll(ctx context.Context) ([]entities.Meaning, error) {

	var meanings []entities.Meaning
	if srv.cache.Get(ctx, CACHE_GET_ALL_MEANINGS, &meanings) {
		return meanings, nil
	}

	meanings, err := srv.repositories.Meaning.GetAll(ctx)
	if err != nil {
		srv.log.ErrorContext(ctx, "Meaning.Service.GetAll", err)
		return nil, err
	}

	srv.cache.Set(ctx, CACHE_GET_ALL_MEANINGS, meanings)

	return meanings, nil
}

func (srv *services) GetByID(ctx context.Context, ID uint64) (*entities.Meaning, error) {

	meaningWanted, err := srv.repositories.Meaning.GetByID(ctx, ID)
	if err != nil {
		srv.log.ErrorContext(ctx, "Meaning.Service.GetByID", ID, err)
		return nil, err
	}

	return meaningWanted, nil
}
