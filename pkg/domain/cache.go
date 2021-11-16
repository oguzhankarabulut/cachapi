package domain

import (
	"errors"
	"sync"
)

var (
	cache = make(map[string]interface{})
	mutex sync.Mutex
	nonExistingKeyError = errors.New("key is not existed")
	nilValueError       = errors.New("value can not be empty")
)

type CacheService interface {
	Read() error
	Write() error
}

type Cache struct {
	key string
	value interface{}
}


// NewCache create new cache instance
func NewCache(k string, v interface{}) *Cache {
	return &Cache{
		key: k,
		value: v,
	}
}

// Set set cache to memory
func Set(c *Cache) error {
	if c.value == nil {
		return nilValueError
	}
	mutex.Lock()
	cache[c.key] = c.value
	mutex.Unlock()
	return nil
}

func Get(k string) (*Cache, error) {
	mutex.Lock()
	v, ok := cache[k]
	mutex.Unlock()
	if !ok {
		return nil, nonExistingKeyError
	}
	return NewCache(k, v), nil
}

func All() map[string]interface{} {
	return cache
}

func AllP() *map[string]interface{} {
	return &cache
}

func Flush() {
	cache = make(map[string]interface{})
}

func(c *Cache) GetKey() string {
	return c.key
}

func(c *Cache) GetValue() interface{} {
	return c.value
}
