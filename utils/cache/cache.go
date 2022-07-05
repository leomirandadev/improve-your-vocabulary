//go:generate mockgen -source cache.go -destination mocks/cache_mock.go -package mocks
package cache

import (
	"context"
	"time"
)

// Cache é a interface do pacote de cache
type Cache interface {
	Get(ctx context.Context, key string, v interface{}) bool
	Set(ctx context.Context, key string, v interface{}) bool

	Del(ctx context.Context, key string) bool

	WithExpiration(d time.Duration) Cache
}

// Options struct de opções para a criação de uma instancia do cache
type Options struct {
	Expiration time.Duration
	URL        string
}
