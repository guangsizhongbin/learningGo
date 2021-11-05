package test

import (
	"sync"
	"testing"
)

var (
	data   = make(map[string]string)
	locker sync.RWMutex
)

func WriteToMap(k, v string) {
	locker.Lock()
	defer locker.Unlock()
	data[k] = v
}

func ReadFromMap(k string) string {
	locker.RLock()
	defer locker.RUnlock()
	return data[k]
}

var pairs = []struct {
	k string
	v string
}{
	{"polaris", "徐新华"},
	{"studygolang", "GO 语言中文网"},
	{"polaris", "徐新华 1"},
	{"stdlib", "Go 语言标准库"},
	{"polaris1", "徐新华"},
	{"studygolang1", "GO 语言中文网"},
	{"polaris1", "徐新华 1"},
	{"stdlib1", "Go 语言标准库"},
}

func TestWriteToMap(t *testing.T) {
	t.Parallel()
	for _, tt := range pairs {
		WriteToMap(tt.k, tt.v)
	}
}

func TestReadFromMap(t *testing.T) {
	t.Parallel()
	for _, tt := range pairs {
		actual := ReadFromMap(tt.k)
		if actual != tt.v {
			t.Errorf("the value of key(%s) is %s, expected: %s", tt.k, actual, tt.v)
		}
	}
}
