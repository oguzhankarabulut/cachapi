package infrastructure

import (
	"cachapi/pkg/domain"
	"fmt"
	"os"
	"testing"
)

var a = NewCacheRepository()

func Test_SaveLocal(t *testing.T) {
	err := a.saveLocal([]byte{123, 34, 102, 111, 111, 34, 58, 34, 98, 97, 114, 34, 125})
	if err != nil {
		t.Errorf("save error: %v", err)
	}
	c := domain.All()
	for k, v := range c {
		if k != "foo" {
			t.Error("key error")
		}
		if v != "bar" {
			t.Error("value error")
		}
	}
}

func Test_ReadFile(t *testing.T) {
	n := now()
	f, err := os.Create(fmt.Sprintf("%s%d%s", folder, n, fileSuffix))
	if err != nil {
		t.Errorf("create error: %v", err)
	}
	defer f.Close()
	if _, err := a.readFile(fmt.Sprintf("%d%s", n, fileSuffix)); err != nil {
		t.Errorf("read file error: %v", err)
	}
}

func Test_LastSavedFile(t *testing.T) {
	n := now()
	f, err := os.Create(fmt.Sprintf("%s%d%s", folder, n, fileSuffix))
	if err != nil {
		t.Errorf("create error: %v", err)
	}
	defer f.Close()

	lf, err := a.lastSavedFile()
	if  err != nil {
		t.Errorf("last saved file error: %v", err)
	}
	if lf != fmt.Sprintf("%d%s", n, fileSuffix) {
		t.Error("file name wrong")
	}
}

func Test_Write(t *testing.T) {
	nc := domain.NewCache("foo", "bar")
	if err := domain.Set(nc); err != nil {
		t.Errorf("set error: %v", err)
	}
	err := a.Write(); if err != nil {
		t.Errorf("write error: %v", err)
	}
}

func Test_Read(t *testing.T) {
	if err := a.Read(); err != nil {
		t.Errorf("read error: %v", err)
	}
	c := domain.All()
	if c["foo"] != "bar" {
		t.Error("unsuccessful read operation")
	}

}
