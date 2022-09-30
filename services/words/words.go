package words

import (
	"context"

	"github.com/leomirandadev/improve-your-vocabulary/entities"
	"github.com/leomirandadev/improve-your-vocabulary/repositories"
	"github.com/leomirandadev/improve-your-vocabulary/utils/logger"
	"github.com/leomirandadev/improve-your-vocabulary/utils/tracer"
)

type IService interface {
	Create(ctx context.Context, newWord entities.WordRequest) (*entities.Word, error)
	GetByID(ctx context.Context, ID uint64, ownerID uint64) (*entities.Word, error)
	GetAll(ctx context.Context, ownerID uint64) ([]entities.Word, error)
}

type services struct {
	repositories *repositories.Container
	log          logger.Logger
}

func New(repo *repositories.Container, log logger.Logger) IService {
	return &services{repositories: repo, log: log}
}

func (srv *services) Create(ctx context.Context, newWord entities.WordRequest) (*entities.Word, error) {
	ctx, tr := tracer.Span(ctx, "services.words.create")
	defer tr.End()

	id, err := srv.repositories.Database.Word.Create(ctx, newWord)
	if err != nil {
		srv.log.ErrorContext(ctx, "Word.Service.Create", err)
		return nil, err
	}

	wordCreated, err := srv.repositories.Database.Word.GetByID(ctx, id, newWord.UserID)
	if err != nil {
		srv.log.ErrorContext(ctx, "Word.Service.GetByID", err)
		return nil, err
	}

	srv.repositories.Cache.Word.DeleteAll(ctx)

	return wordCreated, nil
}

func (srv *services) GetAll(ctx context.Context, ownerID uint64) ([]entities.Word, error) {

	ctx, tr := tracer.Span(ctx, "services.words.get_all")
	defer tr.End()

	if words, err := srv.repositories.Cache.Word.GetAll(ctx, ownerID); err == nil {
		srv.log.ErrorContext(ctx, "Word.Service.cache.GetAll", err)
		return words, nil
	}

	words, err := srv.repositories.Database.Word.GetAll(ctx, ownerID)
	if err != nil {
		srv.log.ErrorContext(ctx, "Word.Service.sql.GetAll", err)
		return nil, err
	}

	for i := range words {
		srv.fillMeanings(ctx, &words[i])
	}

	srv.repositories.Cache.Word.SetAll(ctx, words)

	return words, nil
}

func (srv *services) GetByID(ctx context.Context, ID uint64, ownerID uint64) (*entities.Word, error) {
	ctx, tr := tracer.Span(ctx, "services.words.get_by_id")
	defer tr.End()

	wordWanted, err := srv.repositories.Database.Word.GetByID(ctx, ID, ownerID)
	if err != nil {
		srv.log.ErrorContext(ctx, "Word.Service.GetByID", ID, err)
		return nil, err
	}

	srv.fillMeanings(ctx, wordWanted)

	return wordWanted, nil
}

func (srv *services) fillMeanings(ctx context.Context, word *entities.Word) {
	ctx, tr := tracer.Span(ctx, "services.words.fill_meanings")
	defer tr.End()

	meanings, err := srv.repositories.Database.Meaning.GetByWordID(ctx, word.ID)
	if err != nil {
		srv.log.ErrorContext(ctx, "fillMeanings")
	}
	word.Meanings = meanings
}
