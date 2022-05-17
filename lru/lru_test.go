package lru_test

import (
	"testing"

	Lru "github.com/haobinfei/geeCache/lru"
)

type String string

func (d String) Len() int {
	return len(d)
}

func TestGet(t *testing.T) {
	lru := Lru.New(int64(0), nil)
	lru.Add("key1", String("1234"))

	if _, ok := lru.Get("key1"); ok {
		t.Fatalf("cache")
	}
}
