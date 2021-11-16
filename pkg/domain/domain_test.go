package domain

import (
	"testing"
)

func Test_NewCache(t *testing.T) {
	var v interface{}
	v = "value"
	nc := NewCache("key", v)
	if nc.key != "key" {
		t.Error("key set error")
	}
	if nc.value != "value" {
		t.Error("value set error")
	}
}

func Test_SetWithExistValue(t *testing.T) {
	nc := NewCache("key", "value")
	if err := Set(nc); err != nil {
		t.Errorf("set error: %v", err)
	}
}

func Test_SetWithNilValue(t *testing.T) {
	var v interface{}
	nc := NewCache("key", v)
	if err := Set(nc); err == nil {
		t.Errorf("set error: %v", err)
	}
}

func Test_GetWithExistKey(t *testing.T) {
	nc := NewCache("key", "value")
	if err := Set(nc); err != nil {
		t.Errorf("set error: %v", err)
	}
	c, err := Get("key")
	if err != nil {
		t.Errorf("get value exist error: %v", err)
	}
	if c.value != "value" {
		t.Error("value is not equal")
	}
}

func Test_GetWithNonExistKey(t *testing.T) {
	if _, err := Get("nonExistKey"); err == nil {
		t.Errorf("get value exist error: %v", err)
	}
}

func Test_All(t *testing.T) {
	nc := NewCache("key", "value")
	if err := Set(nc); err != nil {
		t.Errorf("set error: %v", err)
	}
	a := All()
	if a["key"] != "value" {
		t.Error("all error")
	}
}

func Test_Flush(t *testing.T) {
	nc := NewCache("key", "value")
	if err := Set(nc); err != nil {
		t.Errorf("set error: %v", err)
	}
	Flush()
	if len(All()) != 0 {
		t.Error("flush error")
	}
}

func Test_GetKey(t *testing.T) {
	nc := NewCache("key", "value")
	if nc.GetKey() != "key" {
		t.Error("get error")
	}
}

func Test_GetValue(t *testing.T) {
	nc := NewCache("key", "value")
	if nc.GetValue() != "value" {
		t.Error("get error")
	}
}