package meanings

import (
	"context"
	"errors"

	"github.com/leomirandadev/improve-your-vocabulary/internal/entities"
	cacheLib "github.com/leomirandadev/improve-your-vocabulary/pkg/cache"
	"github.com/leomirandadev/improve-your-vocabulary/pkg/tracer"
)

type meaningsImpl struct {
	cache cacheLib.Cache
}

func NewCache(cache cacheLib.Cache) ICache {
	return &meaningsImpl{cache: cache}
}

func (c *meaningsImpl) GetAll(ctx context.Context) ([]entities.Meaning, error) {
	ctx, tr := tracer.Span(ctx, "repositories.cache.meanings.get_all")
	defer tr.End()

	var meanings []entities.Meaning
	if c.cache.Get(ctx, CACHE_GET_ALL_MEANINGS, &meanings) {
		return meanings, nil
	}

	return meanings, errors.New("not found on cache")
}

func (c *meaningsImpl) DeleteAll(ctx context.Context) error {
	ctx, tr := tracer.Span(ctx, "repositories.cache.meanings.delete_all")
	defer tr.End()

	if c.cache.Del(ctx, CACHE_GET_ALL_MEANINGS) {
		return nil
	}
	return errors.New("we can't delete from cache")
}

func (c *meaningsImpl) SetAll(ctx context.Context, meanings []entities.Meaning) error {
	ctx, tr := tracer.Span(ctx, "repositories.cache.meanings.set_all")
	defer tr.End()

	if c.cache.WithExpiration(CACHE_GET_ALL_MEANINGS_EXP).Set(ctx, CACHE_GET_ALL_MEANINGS, meanings) {
		return nil
	}
	return errors.New("we can't set on cache")
}
