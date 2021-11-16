package main

import (
	"cachapi/pkg/api"
	"cachapi/pkg/infrastructure"
	"cachapi/pkg/worker"
	"log"
)

func main() {
	cr := infrastructure.NewCacheRepository()
	if err := cr.Read(); err != nil {
		if err == infrastructure.FileNotExistErr {
			log.Printf("read error: %v", err)
		} else {
			panic(err)
		}
	}
	cw := worker.NewCacheWorker(cr)
	go cw.Run()
	api.Server()
}
