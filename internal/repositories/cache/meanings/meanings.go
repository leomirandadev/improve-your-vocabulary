//go:generate mockgen -source meanings.go -destination mocks/meanings_mock.go -package mocks
package meanings

import (
	"context"

	"github.com/leomirandadev/improve-your-vocabulary/internal/entities"
)

type ICache interface {
	GetAll(ctx context.Context) ([]entities.Meaning, error)
	DeleteAll(ctx context.Context) error
	SetAll(ctx context.Context, meanings []entities.Meaning) error
}
