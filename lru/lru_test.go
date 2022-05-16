package lru_test

import (
	"testing"

	Lru "github.com/geeCache/lru"
)

type String string

func (d String) Len() int {
	return len(d)
}

func TestGet(t *testing.T) {
	lru := Lru.New(int64(0), nil)

}
