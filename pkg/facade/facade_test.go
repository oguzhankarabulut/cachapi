package facade

import (
	"cachapi/pkg/domain"
	"testing"
)

var a = NewCacheFacade()

func Test_SetWithExistValue(t *testing.T) {
	nc, err := a.Set("key", "value")
	if err != nil {
		t.Errorf("set error: %v", err)
	}
	if nc.GetKey() != "key" {
		t.Error("set key error")
	}
	if nc.GetValue() != "value" {
		t.Error("set value error")
	}
}

func Test_SetWithNilValue(t *testing.T) {
	var v interface{}
	if _, err := a.Set("key", v); err == nil {
		t.Errorf("set error: %v", err)
	}
}

func Test_GetWithExistKey(t *testing.T) {
	if _, err := a.Set("key", "value"); err != nil {
		t.Errorf("set error: %v", err)
	}
	if _, err := a.Get("key"); err != nil {
		t.Errorf("get error: %v", err)
	}
}

func Test_GetWithNonExistKey(t *testing.T) {
	if _, err := a.Set("key", "value"); err != nil {
		t.Errorf("set error: %v", err)
	}
	if _, err := a.Get("nonExistKey"); err == nil {
		t.Errorf("get error: %v", err)
	}
}

func Test_FlushCache(t *testing.T) {
	if _, err := a.Set("key", "value"); err != nil {
		t.Errorf("set error: %v", err)
	}
	a.FlushCache()
	if len(domain.All()) > 0 {
		t.Error("flush cache error")
	}
}

