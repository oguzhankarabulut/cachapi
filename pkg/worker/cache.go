package worker

import (
	"cachapi/pkg/domain"
	"fmt"
	"log"
	"time"
)

type cacheWorker struct {
	cs domain.CacheService
}

func NewCacheWorker(cacheService domain.CacheService) *cacheWorker {
	return &cacheWorker{
		cs: cacheService,
	}
}

func (cw cacheWorker) Run() {
	for {
		time.Sleep(30 * time.Minute)
		if len(domain.All()) != 0 {
			if err := cw.cs.Write(); err != nil {
				log.Println(fmt.Sprintf("Worker Error: %v", err))
			}
		}
	}
}
