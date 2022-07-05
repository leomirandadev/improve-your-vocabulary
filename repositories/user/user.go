package user

import (
	"context"

	"github.com/leomirandadev/improve-your-vocabulary/entities"
)

type IRepository interface {
	Create(ctx context.Context, newUser entities.User) (uint64, error)
	GetUserByEmail(ctx context.Context, userLogin entities.UserAuth) (entities.User, error)
	GetByID(ctx context.Context, ID uint64) (entities.UserResponse, error)
}
