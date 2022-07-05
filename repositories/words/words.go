package words

import (
	"context"

	"github.com/leomirandadev/improve-your-vocabulary/entities"
)

type IRepository interface {
	Create(ctx context.Context, newWord entities.WordRequest) (uint64, error)
	GetByID(ctx context.Context, ID, ownerID uint64) (*entities.Word, error)
	GetAll(ctx context.Context, ownerID uint64) ([]entities.Word, error)
}
