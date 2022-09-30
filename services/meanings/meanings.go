package meanings

import (
	"context"

	"github.com/leomirandadev/improve-your-vocabulary/entities"
	"github.com/leomirandadev/improve-your-vocabulary/repositories"
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
}

func New(repo *repositories.Container, log logger.Logger) IService {
	return &services{repositories: repo, log: log}
}

func (srv *services) Create(ctx context.Context, newMeaning entities.MeaningRequest) (*entities.Meaning, error) {
	id, err := srv.repositories.Database.Meaning.Create(ctx, newMeaning)
	if err != nil {
		srv.log.ErrorContext(ctx, "Meaning.Service.sql.Create", err)
		return nil, err
	}

	meaningCreated, err := srv.repositories.Database.Meaning.GetByID(ctx, id)
	if err != nil {
		srv.log.ErrorContext(ctx, "Meaning.Service.sql.GetByID", err)
		return nil, err
	}

	srv.repositories.Cache.Meaning.DeleteAll(ctx)
	srv.repositories.Cache.Word.DeleteAll(ctx)

	return meaningCreated, nil
}

func (srv *services) GetAll(ctx context.Context) ([]entities.Meaning, error) {

	if meanings, err := srv.repositories.Cache.Meaning.GetAll(ctx); err == nil {
		srv.log.ErrorContext(ctx, "Meaning.Service.cache.GetAll", err)
		return meanings, nil
	}

	meanings, err := srv.repositories.Database.Meaning.GetAll(ctx)
	if err != nil {
		srv.log.ErrorContext(ctx, "Meaning.Service.sql.GetAll", err)
		return nil, err
	}

	srv.repositories.Cache.Meaning.SetAll(ctx, meanings)

	return meanings, nil
}

func (srv *services) GetByID(ctx context.Context, ID uint64) (*entities.Meaning, error) {

	meaningWanted, err := srv.repositories.Database.Meaning.GetByID(ctx, ID)
	if err != nil {
		srv.log.ErrorContext(ctx, "Meaning.Service.sql.GetByID", ID, err)
		return nil, err
	}

	return meaningWanted, nil
}
