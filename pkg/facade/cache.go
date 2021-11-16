package facade

import "cachapi/pkg/domain"

type cacheFacade struct {

}

type CacheFacade interface {
	Get(k string) (*domain.Cache, error)
	Set(k string, v interface{}) (*domain.Cache, error)
	FlushCache()
}

func NewCacheFacade() *cacheFacade {
	return &cacheFacade{}
}

func (c cacheFacade) Get(k string) (*domain.Cache, error) {
	nc, err := domain.Get(k)
	if err != nil {
		return nil, err
	}
	return nc, nil
}

func (c cacheFacade) Set(k string, v interface{}) (*domain.Cache, error) {
	nc := domain.NewCache(k, v)
	if err := domain.Set(nc); err != nil {
		return nil, err
	}
	return nc, nil
}

func (c cacheFacade) FlushCache() {
	domain.Flush()
}


