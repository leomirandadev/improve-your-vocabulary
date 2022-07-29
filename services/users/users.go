package users

import (
	"context"

	"github.com/leomirandadev/improve-your-vocabulary/entities"
	"github.com/leomirandadev/improve-your-vocabulary/repositories"
	"github.com/leomirandadev/improve-your-vocabulary/utils/hasher"
	"github.com/leomirandadev/improve-your-vocabulary/utils/logger"
)

type IService interface {
	Create(ctx context.Context, newUser entities.UserRequest) error
	GetByID(ctx context.Context, ID uint64) (entities.UserResponse, error)
	GetUserByLogin(ctx context.Context, userLogin entities.UserAuth) (entities.UserResponse, error)
}

type services struct {
	repositories *repositories.Container
	log          logger.Logger
}

func New(repo *repositories.Container, log logger.Logger) IService {
	return &services{repositories: repo, log: log}
}

func (srv *services) Create(ctx context.Context, newUser entities.UserRequest) error {
	hasherBcrypt := hasher.NewBcryptHasher()
	passwordHashed, errHash := hasherBcrypt.Generate(newUser.Password)

	if errHash != nil {
		srv.log.Error("Srv.Create: ", "Error generate hash password: ", newUser)
		return errHash
	}

	newUser.Password = passwordHashed
	_, err := srv.repositories.Sql.User.Create(ctx, newUser)

	return err
}

func (srv *services) GetUserByLogin(ctx context.Context, userLogin entities.UserAuth) (entities.UserResponse, error) {

	userFound, err := srv.repositories.Sql.User.GetUserByEmail(ctx, userLogin)

	if err != nil {
		srv.log.Error("Srv.Auth: ", "User not found", userFound)
		return entities.UserResponse{}, err
	}

	hasherBcrypt := hasher.NewBcryptHasher()
	err = hasherBcrypt.Compare(userFound.Password, userLogin.Password)
	if err != nil {
		return entities.UserResponse{}, err
	}

	return entities.UserResponse{
		ID:        userFound.ID,
		NickName:  userFound.NickName,
		Name:      userFound.Name,
		Email:     userFound.Email,
		Role:      userFound.Role,
		CreatedAt: userFound.CreatedAt,
		UpdatedAt: userFound.UpdatedAt,
	}, nil
}

func (srv *services) GetByID(ctx context.Context, ID uint64) (entities.UserResponse, error) {
	return srv.repositories.Sql.User.GetByID(ctx, ID)
}
