package cache

import (
	"time"
)

// Fetcher can fetch items that will be held by Cache.
type Fetcher[T any] interface {
	Fetch() (T, error)
}

// Cache contains cached items and keeps track of time.
type Cache[T any] struct {
	Fetcher[T]
	cache            T
	lastRefreshTime  time.Time
	expirationPeriod time.Duration
}

func New[T any](fetcher Fetcher[T], expirationPeriod time.Duration) *Cache[T] {
	return &Cache[T]{
		Fetcher: fetcher,
		// Cache is expired by default as there are no items in it.
		lastRefreshTime:  time.Now().Add(-2 * expirationPeriod),
		expirationPeriod: expirationPeriod,
	}
}

// Get cached item(s) that are refreshed if expired or returned as is.
func (c *Cache[T]) Get() (T, error) {
	err := c.refresh()
	return c.cache, err
}

// refresh expired cache if necessary.
func (c *Cache[T]) refresh() (err error) {
	if c.isFresh() {
		return
	}

	c.cache, err = c.Fetch()
	if err != nil {
		return
	}

	c.lastRefreshTime = time.Now()
	return
}

// isFresh tells us whether it's time to refresh our cache.
func (c *Cache[T]) isFresh() bool {
	return time.Now().Sub(c.lastRefreshTime) < c.expirationPeriod
}
