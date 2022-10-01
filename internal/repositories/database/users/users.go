package users

import (
	"context"

	"github.com/leomirandadev/improve-your-vocabulary/internal/entities"
)

type IRepository interface {
	Create(ctx context.Context, newUser entities.UserRequest) (uint64, error)
	GetUserByEmail(ctx context.Context, userLogin entities.UserAuth) (entities.User, error)
	GetByID(ctx context.Context, ID uint64) (entities.UserResponse, error)
}
