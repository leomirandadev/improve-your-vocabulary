package words

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

func (repo *repoSqlx) Create(ctx context.Context, newWord entities.WordRequest) (uint64, error) {

	result, err := repo.writer.ExecContext(ctx, `
		INSERT INTO words (word) VALUES (:word)
	`, newWord)

	if err != nil {
		repo.log.ErrorContext(ctx, "Word.SqlxRepo.Create", err)
		return 0, errors.New("Não foi possível criar palavra")
	}

	id, err := result.LastInsertId()
	if err != nil {
		repo.log.ErrorContext(ctx, "Word.SqlxRepo.LastInsertId", err)
		return 0, errors.New("Não foi possível criar o usuário")
	}

	return uint64(id), nil
}

func (repo *repoSqlx) GetByID(ctx context.Context, ID uint64) (*entities.Word, error) {

	var word entities.Word

	err := repo.reader.GetContext(ctx, &word, `
		SELECT id, word, created_at, updated_at FROM words WHERE id = ?
	`, ID)

	if err != nil {
		repo.log.ErrorContext(ctx, "Word.SqlxRepo.GetByID", "Error on get Word ID: ", ID, err)
		return nil, errors.New("Palavra não encontrado")
	}

	return &word, nil
}

func (repo *repoSqlx) GetAll(ctx context.Context) ([]entities.Word, error) {

	words := make([]entities.Word, 0)

	err := repo.reader.SelectContext(ctx, &words, `
		SELECT id, word, created_at, updated_at FROM words
	`)

	if err != nil {
		repo.log.ErrorContext(ctx, "Word.SqlxRepo.GetAll", err)
		return words, errors.New("Nenhuma palavra encontrada")
	}

	return words, nil
}
