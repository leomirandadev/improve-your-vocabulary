package words

import (
	"context"
	"errors"

	"github.com/leomirandadev/improve-your-vocabulary/entities"
	cacheLib "github.com/leomirandadev/improve-your-vocabulary/utils/cache"
	"github.com/leomirandadev/improve-your-vocabulary/utils/tracer"
)

type wordsImpl struct {
	cache cacheLib.Cache
}

func NewCache(cache cacheLib.Cache) ICache {
	return &wordsImpl{cache: cache}
}

func (c *wordsImpl) GetAll(ctx context.Context, ownerID uint64) ([]entities.Word, error) {
	ctx, tr := tracer.Span(ctx, "repositories.cache.words.get_all")
	defer tr.End()

	var words []entities.Word
	if c.cache.Get(ctx, CACHE_GET_ALL_WORDS, &words) {
		return words, nil
	}

	return words, errors.New("not found on cache")
}

func (c *wordsImpl) DeleteAll(ctx context.Context) error {
	ctx, tr := tracer.Span(ctx, "repositories.cache.words.delete_all")
	defer tr.End()

	if c.cache.Del(ctx, CACHE_GET_ALL_WORDS) {
		return nil
	}
	return errors.New("we can't delete from cache")
}

func (c *wordsImpl) SetAll(ctx context.Context, words []entities.Word) error {
	ctx, tr := tracer.Span(ctx, "repositories.cache.words.set_all")
	defer tr.End()

	if c.cache.WithExpiration(CACHE_GET_ALL_WORDS_EXP).Set(ctx, CACHE_GET_ALL_WORDS, words) {
		return nil
	}

	return errors.New("we can't set on cache")
}
