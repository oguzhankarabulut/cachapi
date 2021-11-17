package domain

import (
	"errors"
	"sync"
)

var (
	cache = make(map[string]interface{})
	mutex sync.Mutex
	nonExistingKeyError = errors.New("key is not existed")
)

type CacheService interface {
	Read() error
	Write() error
}

type Cache struct {
	key string
	value interface{}
}

func NewCache(k string, v interface{}) *Cache {
	return &Cache{
		key: k,
		value: v,
	}
}

// Set set key-value to map. It has locking mechanism to avoid concurrent read/write error
func Set(c *Cache) {
	mutex.Lock()
	cache[c.key] = c.value
	mutex.Unlock()
	return
}

// Get get value by key from map. It has locking mechanism to avoid concurrent read/write error
func Get(k string) (*Cache, error) {
	mutex.Lock()
	v, ok := cache[k]
	mutex.Unlock()
	if !ok {
		return nil, nonExistingKeyError
	}
	return NewCache(k, v), nil
}

// All return the map which is used for memory store
func All() map[string]interface{} {
	return cache
}

// AllP return the pointer type of map which is used for memory store
func AllP() *map[string]interface{} {
	return &cache
}

// Flush flush the map
func Flush() {
	cache = make(map[string]interface{})
}

func(c *Cache) GetKey() string {
	return c.key
}

func(c *Cache) GetValue() interface{} {
	return c.value
}
