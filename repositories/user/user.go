package user

import (
	"context"

	"github.com/leomirandadev/improve-your-vocabulary/entities"
)

type UserRepository interface {
	Create(ctx context.Context, newUser entities.User) error
	GetUserByEmail(ctx context.Context, userLogin entities.UserAuth) (entities.User, error)
	GetByID(ctx context.Context, ID int64) (entities.UserResponse, error)
}
