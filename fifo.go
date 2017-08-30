package caches

import "container/list"

// NewFIFO ...
func NewFIFO(maxEntries int) Cache {
	c := &fifo{
		keys:       make(map[interface{}]*list.Element),
		entries:    list.New(),
		maxEntries: maxEntries,
	}
	return c
}

type fifo struct {
	keys       map[interface{}]*list.Element
	entries    *list.List
	maxEntries int
}

// Add key-value to the cache
func (c *fifo) Add(key, value interface{}) (old interface{}) {
	if elem, ok := c.keys[key]; ok {
		old = elem.Value.(*entry).value
		elem.Value.(*entry).value = value
		return old
	}

	elem := c.entries.PushFront(&entry{key, value})
	c.keys[key] = elem
	if c.maxEntries != 0 && c.Size() > c.maxEntries {
		c.Pop()
	}
	return nil
}

// Get value for given key from the cache
func (c *fifo) Get(key interface{}) (value interface{}, ok bool) {
	elem, ok := c.keys[key]
	if !ok {
		return nil, false
	}
	return elem.Value.(*entry).value, true
}

// Pop removes needless element from the cache
func (c *fifo) Pop() (key, value interface{}) {
	elem := c.entries.Back()
	if elem == nil {
		return nil, nil
	}
	key = elem.Value.(*entry).key
	value = elem.Value.(*entry).value
	c.removElement(elem)
	return key, value
}

// Remove key-value from the cache by key
func (c *fifo) Remove(key interface{}) (value interface{}, ok bool) {
	if elem, ok := c.keys[key]; ok {
		c.removElement(elem)
		return elem.Value.(*entry).value, true
	}
	return nil, false
}

// removeElement ...
func (c *fifo) removElement(e *list.Element) {
	key := e.Value.(*entry).key
	c.entries.Remove(e)
	delete(c.keys, key)
}

// Size returns size of the cache
func (c *fifo) Size() int {
	return c.entries.Len()
}

// Clear the cache
func (c *fifo) Clear() {
	c.entries = nil
	c.keys = nil
}
