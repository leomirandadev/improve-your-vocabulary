package meanings

import (
	"context"
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/leomirandadev/improve-your-vocabulary/entities"
	"github.com/leomirandadev/improve-your-vocabulary/utils/logger"
)

type repoSqlx struct {
	log    logger.Logger
	writer *sqlx.DB
	reader *sqlx.DB
}

func NewSqlx(log logger.Logger, writer, reader *sqlx.DB) IRepository {
	return &repoSqlx{log: log, writer: writer, reader: reader}
}

func (repo *repoSqlx) Create(ctx context.Context, newMeaning entities.MeaningRequest) (uint64, error) {

	result, err := repo.writer.ExecContext(ctx, `
		INSERT INTO meanings (meaning, word_id) VALUES (?, ?)
	`, newMeaning.Meaning, newMeaning.WordID)

	if err != nil {
		repo.log.ErrorContext(ctx, "Meaning.SqlxRepo.Create", err)
		return 0, errors.New("Não foi possível criar o significado")
	}

	id, err := result.LastInsertId()
	if err != nil {
		repo.log.ErrorContext(ctx, "Meaning.SqlxRepo.LastInsertId", err)
		return 0, errors.New("Não foi possível criar o significado")
	}

	return uint64(id), nil
}

func (repo *repoSqlx) GetByID(ctx context.Context, ID uint64) (*entities.Meaning, error) {

	var meaning entities.Meaning

	err := repo.reader.GetContext(ctx, &meaning, `
		SELECT id, meaning, word_id, created_at, updated_at FROM meanings WHERE id = ?
	`, ID)

	if err != nil {
		repo.log.ErrorContext(ctx, "Meaning.SqlxRepo.GetByID", "Error on get meaning ID: ", ID, err)
		return nil, errors.New("Significado não encontrado")
	}

	return &meaning, nil
}

func (repo *repoSqlx) GetAll(ctx context.Context) ([]entities.Meaning, error) {

	meanings := make([]entities.Meaning, 0)

	err := repo.reader.SelectContext(ctx, &meanings, `
		SELECT id, meaning, word_id, created_at, updated_at FROM meanings
	`)

	if err != nil {
		repo.log.ErrorContext(ctx, "Meaning.SqlxRepo.GetAll", err)
		return meanings, errors.New("Nenhum significado encontrado")
	}

	return meanings, nil
}

func (repo *repoSqlx) GetByWordID(ctx context.Context, wordID uint64) ([]entities.Meaning, error) {

	meanings := make([]entities.Meaning, 0)

	err := repo.reader.SelectContext(ctx, &meanings, `
		SELECT 
			id,
			meaning,
			word_id,
			created_at,
			updated_at
		FROM meanings
		WHERE word_id = ?
	`, wordID)

	if err != nil {
		repo.log.ErrorContext(ctx, "Meaning.SqlxRepo.GetByWordID", wordID, err)
		return meanings, errors.New("Nenhum significado encontrado")
	}

	return meanings, nil
}
