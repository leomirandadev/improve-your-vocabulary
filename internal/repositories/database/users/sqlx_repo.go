package users

import (
	"context"
	"errors"

	"github.com/jmoiron/sqlx"
	"github.com/leomirandadev/improve-your-vocabulary/internal/entities"
	"github.com/leomirandadev/improve-your-vocabulary/pkg/logger"
	"github.com/leomirandadev/improve-your-vocabulary/pkg/tracer"
)

type repoSqlx struct {
	log    logger.Logger
	writer *sqlx.DB
	reader *sqlx.DB
}

func NewSqlx(log logger.Logger, writer, reader *sqlx.DB) IRepository {
	return &repoSqlx{log: log, writer: writer, reader: reader}
}

func (repo *repoSqlx) Create(ctx context.Context, newUser entities.UserRequest) (uint64, error) {
	ctx, tr := tracer.Span(ctx, "repositories.database.users.create")
	defer tr.End()

	result, err := repo.writer.ExecContext(ctx, `
		INSERT INTO users (nick_name,name,email,password,role) VALUES (?, ?, ?, ?, ?)
	`, newUser.NickName, newUser.Name, newUser.Email, newUser.Password, newUser.Role)

	if err != nil {
		repo.log.ErrorContext(ctx, "User.SqlxRepo.Create", err)
		return 0, errors.New("Não foi possível criar o usuário")
	}

	id, err := result.LastInsertId()
	if err != nil {
		repo.log.ErrorContext(ctx, "Meaning.SqlxRepo.LastInsertId", err)
		return 0, errors.New("Não foi possível criar o usuário")
	}

	return uint64(id), nil
}

func (repo *repoSqlx) GetByID(ctx context.Context, ID uint64) (entities.UserResponse, error) {
	ctx, tr := tracer.Span(ctx, "repositories.database.users.get_by_id")
	defer tr.End()

	var user entities.UserResponse

	err := repo.reader.GetContext(ctx, &user, `
		SELECT 
			id,
			nick_name,
			name,
			email,
			role,
			created_at,
			updated_at
		FROM users 
		WHERE id = ?
		LIMIT 1
	`, ID)

	if err != nil {
		repo.log.ErrorContext(ctx, "User.SqlxRepo.GetByID", "Error on get User ID: ", ID, err)
		return user, errors.New("Usuário não encontrado")
	}

	return user, nil

}

func (repo *repoSqlx) GetUserByEmail(ctx context.Context, userLogin entities.UserAuth) (entities.User, error) {
	ctx, tr := tracer.Span(ctx, "repositories.database.users.get_user_by_email")
	defer tr.End()

	var user entities.User

	err := repo.reader.GetContext(ctx, &user, `
		SELECT 
			id,
			nick_name,
			name,
			email,
			role,
			password,
			created_at,
			updated_at
		FROM users 
		WHERE email = ?
		LIMIT 1
	`, userLogin.Email)

	if err != nil {
		repo.log.ErrorContext(ctx, "User.SqlxRepo.GetUserByEmail", "Error on get User: ", userLogin, err)
		return user, errors.New("Usuário não encontrado")
	}

	return user, nil
}
