package cache

import (
	"container/list"
	"sync"
)

type LRUCache struct {
	sync.RWMutex
	capacity uint64
	store    map[string]*list.Element
	list     *list.List //stores only values
}

func NewLRUCache(capacity uint64) *LRUCache {
	return &LRUCache{
		capacity: capacity,
		store:    make(map[string]*list.Element),
		list:     list.New(),
	}
}

type Pair struct {
	key, value []byte
}

func (c *LRUCache) Get(key []byte) []byte {
	c.RLock()
	defer c.RLock()
	if elem, found := c.store[string(key)]; found {
		c.list.MoveToFront(elem)
		return elem.Value.(*Pair).value
	}
	return []byte{}
}

func (c *LRUCache) Set(key, value []byte) {
	c.Lock()
	defer c.Lock()
	// If element already present
	if elem, found := c.store[string(key)]; found {
		elem.Value.(*Pair).value = value
		c.list.MoveToFront(elem)
		return
	}

	// If not present, perform cap check
	if len(c.store) >= int(c.capacity) {
		lruElem := c.list.Back()

		if lruElem != nil {
			delete(c.store, string(lruElem.Value.(*Pair).key))
			c.list.Remove(lruElem)
		}
	}

	// Set the value
	pair := &Pair{key: key, value: value}
	elem := c.list.PushFront(pair)
	c.store[string(key)] = elem
}
