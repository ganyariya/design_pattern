package go_strategy

import (
	"reflect"
	"testing"
)

func TestCache(t *testing.T) {
	c := NewCache[int]()
	for i := 0; i < 5; i++ {
		c.Push(i)
	}
	if len(c.data) != 5 {
		t.Fatal()
	}

	c.SetStrategy(&FIFO[int]{})
	ret := c.GetElement()
	if ret != 0 || !reflect.DeepEqual([]int{1, 2, 3, 4}, c.data) {
		t.Fatal()
	}

	c.SetStrategy(&FILO[int]{})
	ret = c.GetElement()
	if ret != 4 || !reflect.DeepEqual([]int{1, 2, 3}, c.data) {
		t.Fatal()
	}
}
