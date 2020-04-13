package hw04_lru_cache //nolint:golint,stylecheck
import "sync"

type Key string

type Cache interface {
	Set(key Key, value interface{}) bool
	Get(key Key) (interface{}, bool)
	Clear()
}

type lruCache struct {
	mux      sync.Mutex
	capacity int
	queue    List
	items    map[Key]*listItem
}

type cacheItem struct {
	Key   Key
	Value interface{}
}

func (c *lruCache) Set(key Key, value interface{}) bool {
	c.mux.Lock()

	v, found := c.items[key]
	newItem := cacheItem{key, value}
	if found {
		v.Value = newItem
		c.queue.MoveToFront(v)
	} else {
		c.items[key] = &listItem{Value: newItem}
		c.queue.PushFront(newItem)
	}
	if c.queue.Len() > c.capacity {
		backQueueItem := c.queue.Back()
		c.queue.Remove(backQueueItem)
		delete(c.items, backQueueItem.Value.(cacheItem).Key)
	}
	c.mux.Unlock()
	return found
}

func (c *lruCache) Get(key Key) (interface{}, bool) {
	c.mux.Lock()
	defer c.mux.Unlock()

	item, found := c.items[key]
	if found {
		c.queue.MoveToFront(item)
		return item.Value.(cacheItem).Value, true
	}
	return nil, false
}

func (c *lruCache) Clear() {
	c.queue = NewList()
	c.items = make(map[Key]*listItem)
}

func NewCache(capacity int) Cache {
	return &lruCache{capacity: capacity, queue: NewList(), items: map[Key]*listItem{}}
}
