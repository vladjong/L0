package cache

import (
	"errors"
	"time"
)

type Cache struct {
	items map[string]Item
}

type Item struct {
	Value      interface{}
	Expiration int64
}

func New() *Cache {
	items := make(map[string]Item)
	cache := Cache{
		items: items,
	}
	return &cache
}

func (c *Cache) Set(key string, value interface{}) {
	var expiration int64
	c.items[key] = Item{
		Value:      value,
		Expiration: expiration,
	}
}

func (c *Cache) Len() int {
	n := len(c.items)
	return n
}

func (c *Cache) Get(key string) (interface{}, bool) {
	item, found := c.items[key]
	if !found {
		return nil, false
	}
	if item.Expiration > 0 {
		if time.Now().UnixNano() > item.Expiration {
			return nil, false
		}
	}
	return item.Value, true
}

func (c *Cache) Delete(key string) error {
	if _, found := c.items[key]; !found {
		return errors.New("key not found")
	}
	delete(c.items, key)
	return nil
}
