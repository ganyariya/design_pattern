package go_template_method

import (
	"testing"
)

func TestBinary(t *testing.T) {
	var toBinaries []ToBinary = []ToBinary{
		{
			NewMemoryBinary("tmp"),
		},
	}
	bytes := toBinaries[0].Convert()

	expectedBytes := []byte("HELLO")
	for i := 0; i < len(expectedBytes); i++ {
		if bytes[i] != expectedBytes[i] {
			t.Fatalf("error")
		}
	}
}
