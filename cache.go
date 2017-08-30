package caches

// Cache is a common interface for all caches
type Cache interface {
	// Add key-value to the cache, returns oldValue for key or nil
	// Might execute Pop method if number of keys is bigger than maxEntries
	Add(key interface{}, value interface{}) (oldValue interface{})

	// Get value for given key from the cache
	Get(key interface{}) (value interface{}, ok bool)

	// Pop removes needless element from the cache
	Pop() (key, value interface{})

	// Remove key-value from the cache by key
	Remove(key interface{}) (value interface{}, ok bool)

	// Size returns size of the cache
	Size() int

	// Clear the cache
	Clear()
}

// entry is a pair of key-value in caches
type entry struct {
	key   interface{}
	value interface{}
}

type cacheType int

const (
	// FIFO cache, simple queue
	FIFO cacheType = iota
	// LRU cache
	LRU
	// MRU cache
	MRU
)

// New ...
func New(cache cacheType, maxEntries ...int) Cache {
	size := 0
	if len(maxEntries) > 0 {
		size = maxEntries[0]
	}
	switch cache {
	case FIFO:
		return NewFIFO(size)

	case LRU:
		// return NewLRU(size)
		return nil

	case MRU:
		// return NewMRU(size)
		return nil

	default:
		return nil
	}
}
