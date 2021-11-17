package facade

import (
	"cachapi/pkg/domain"
	"testing"
)

var a = NewCacheFacade()

func Test_SetWithExistValue(t *testing.T) {
	nc := a.Set("key", "value")
	if nc.GetKey() != "key" {
		t.Error("set key error")
	}
	if nc.GetValue() != "value" {
		t.Error("set value error")
	}
}

func Test_GetWithExistKey(t *testing.T) {
	_ = a.Set("key", "value")
	if _, err := a.Get("key"); err != nil {
		t.Errorf("get error: %v", err)
	}
}

func Test_GetWithNonExistKey(t *testing.T) {
	_ = a.Set("key", "value")
	if _, err := a.Get("nonExistKey"); err == nil {
		t.Errorf("get error: %v", err)
	}
}

func Test_FlushCache(t *testing.T) {
	_ = a.Set("key", "value")
	a.FlushCache()
	if len(domain.All()) > 0 {
		t.Error("flush cache error")
	}
}

