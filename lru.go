package caches

import "container/list"

// NewLRU ...
func NewLRU(maxEntries int) Cache {
	c := &lru{
		keys:       make(map[interface{}]*list.Element),
		entries:    list.New(),
		maxEntries: maxEntries,
	}
	return c
}

// LRU ...
type lru struct {
	keys       map[interface{}]*list.Element
	entries    *list.List
	maxEntries int
}

// Add key-value to the cache
func (c *lru) Add(key interface{}, value interface{}) (oldValue interface{}) {
	if elem, ok := c.keys[key]; ok {
		c.entries.MoveToFront(elem)
		oldValue = elem.Value.(*entry).value
		elem.Value.(*entry).value = value
		return
	}
	elem := c.entries.PushFront(&entry{key, value})
	c.keys[key] = elem
	if c.maxEntries != 0 && c.entries.Len() > c.maxEntries {
		c.Pop()
	}
	return nil
}

// Get value for given key from the cache
func (c *lru) Get(key interface{}) (value interface{}, ok bool) {
	elem, ok := c.keys[key]
	if !ok {
		return nil, false
	}
	c.entries.MoveToFront(elem)
	return elem.Value.(*entry).value, true
}

// Pop removes needless element from the cache
func (c *lru) Pop() (key, value interface{}) {
	elem := c.entries.Front()
	if elem == nil {
		return nil, nil
	}
	key = elem.Value.(*entry).key
	value = elem.Value.(*entry).value
	c.removeElement(elem)
	return key, value
}

// Remove key-value from the cache by key
func (c *lru) Remove(key interface{}) (value interface{}, ok bool) {
	if elem, ok := c.keys[key]; ok {
		c.removeElement(elem)
		return elem.Value.(*entry).value, true
	}
	return nil, false
}

// removeElement ...
func (c *lru) removeElement(e *list.Element) {
	key := e.Value.(*entry).key
	c.entries.Remove(e)
	delete(c.keys, key)
}

// Size returns size of the cache
func (c *lru) Size() int {
	return c.entries.Len()
}

// Clear the cache
func (c *lru) Clear() {
	c.entries = nil
	c.keys = nil
}
