package caches

import "sync"

// NewConcurrent wraps given cache into a thread safe cache
func NewConcurrent(cache Cache) Cache {
	return &concurrent{
		cache: cache,
	}
}

// concurrent is an internal struct for wrapping
type concurrent struct {
	sync.Mutex
	cache Cache
}

// Add key-value to the cache
func (c *concurrent) Add(key, value interface{}) (oldValue interface{}) {
	c.Lock()
	defer c.Unlock()

	return c.cache.Add(key, value)
}

// Get value for given key from the cache
func (c *concurrent) Get(key interface{}) (value interface{}, ok bool) {
	c.Lock()
	defer c.Unlock()

	return c.cache.Get(key)
}

// Pop removes value from the cache
func (c *concurrent) Pop() (key, value interface{}) {
	c.Lock()
	defer c.Unlock()

	return c.cache.Pop()
}

// Remove key-value from the cache by key
func (c *concurrent) Remove(key interface{}) (value interface{}, ok bool) {
	c.Lock()
	defer c.Unlock()

	return c.cache.Remove(key)
}

// Size returns size of the cache
func (c *concurrent) Size() int {
	c.Lock()
	defer c.Unlock()

	return c.cache.Size()
}

// Clear the cache
func (c *concurrent) Clear() {
	c.Lock()
	defer c.Unlock()

	c.cache.Clear()
}
