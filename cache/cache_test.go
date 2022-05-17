package cache_test

import (
	"reflect"
	"testing"

	cache "github.com/haobinfei/geeCache/cache"
)

func TestGetter(t *testing.T) {
	var f cache.Getter = cache.GetterFunc(func(key string) ([]byte, error) {
		return []byte(key), nil
	})

	expect := []byte("key")
	if v, _ := f.Get("key"); !reflect.DeepEqual(v, expect) {
		t.Errorf("callback failed")
	}
}
