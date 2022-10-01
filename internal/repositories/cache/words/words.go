//go:generate mockgen -source words.go -destination mocks/words_mock.go -package mocks
package words

import (
	"context"

	"github.com/leomirandadev/improve-your-vocabulary/internal/entities"
)

type ICache interface {
	GetAll(ctx context.Context, ownerID uint64) ([]entities.Word, error)
	DeleteAll(ctx context.Context) error
	SetAll(ctx context.Context, words []entities.Word) error
}
