package infrastructure

import (
	"cachapi/pkg/domain"
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"strconv"
	"strings"
	"sync"
)

const (
	fileSuffix = "-data.json"
	folder = "/tmp/"
)

var (
	mutex sync.Mutex
	FileNotExistErr = errors.New("no data file")
)

type cacheRepository struct {

}

func NewCacheRepository() *cacheRepository {
	return &cacheRepository{}
}

func (cr cacheRepository) Read() error {
	fn, err := cr.lastSavedFile()
	if err != nil {
		return err
	}

	if fn == "" {
		return FileNotExistErr
	}

	bb, err := cr.readFile(fn)
	if err != nil {
		return err
	}

	if err := cr.saveLocal(bb); err != nil {
		return err
	}
	return nil
}

func (cr cacheRepository) Write() error {
	f, err := os.Create(fmt.Sprintf("%s%d%s", folder, now(), fileSuffix))
	if err != nil {
		return err
	}
	defer f.Close()

	mutex.Lock()
	b, err := json.Marshal(domain.All())
	mutex.Unlock()
	if err != nil {
		return err
	}

	if _, err = f.Write(b); err != nil {
		return err
	}
	return nil
}

func (cr cacheRepository) lastSavedFile() (string, error) {
	ff, err := ioutil.ReadDir(folder)
	if err != nil {
		return "", err
	}

	var lastTime int
	var lastFileName string
	for _, f := range ff {
		tt := strings.Split(f.Name(), fileSuffix)
		if len(tt) > 1 {
			t, err := strconv.Atoi(tt[0])
			if err != nil {
				return "", err
			}
			if t > lastTime {
				lastFileName = f.Name()
			}
		}
	}
	return lastFileName, nil
}

func (cr cacheRepository) readFile(fn string) ([]byte, error) {
	b, err := ioutil.ReadFile(fmt.Sprintf("/tmp/%s", fn))
	if err != nil {
		return nil, err
	}
	return b, nil
}

func (cr cacheRepository) saveLocal(b []byte) error {
	if err := json.Unmarshal(b, domain.AllP()); err != nil {
		return err
	}
	return nil
}
