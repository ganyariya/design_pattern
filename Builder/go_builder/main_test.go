package go_builder

import "testing"

func TestMain(t *testing.T) {
	director := Director{builder: &RedBuilder{}}
	rect := director.Construct()
	if rect.color != "Yellow" {
		t.Fatalf(rect.color)
		t.Fail()
	}
}
