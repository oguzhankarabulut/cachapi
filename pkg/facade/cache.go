package facade

import "cachapi/pkg/domain"

type cacheFacade struct {

}

type CacheFacade interface {
	Get(k string) (*domain.Cache, error)
	Set(k string, v interface{}) *domain.Cache
	FlushCache()
}

func NewCacheFacade() *cacheFacade {
	return &cacheFacade{}
}

// Get get value by key
func (c cacheFacade) Get(k string) (*domain.Cache, error) {
	nc, err := domain.Get(k)
	if err != nil {
		return nil, err
	}
	return nc, nil
}

// Set set key-value with given key and value
func (c cacheFacade) Set(k string, v interface{}) *domain.Cache {
	nc := domain.NewCache(k, v)
	domain.Set(nc)
	return nc
}

// FlushCache executes flush
func (c cacheFacade) FlushCache() {
	domain.Flush()
}


