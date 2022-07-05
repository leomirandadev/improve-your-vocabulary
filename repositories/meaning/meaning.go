package meaning

import (
	"context"

	"github.com/leomirandadev/improve-your-vocabulary/entities"
)

type IRepository interface {
	Create(ctx context.Context, newMeaning entities.MeaningRequest) (uint64, error)
	GetByID(ctx context.Context, ID uint64) (*entities.Meaning, error)
	GetAll(ctx context.Context) ([]entities.Meaning, error)
}
