package go_adapter

import (
	"testing"

	"github.com/ganyariya/design_pattern/Adapter/go_adapter/external"
)

func TestClient(t *testing.T) {
	dollar := 1
	var taxInterface TaxInterface
	taxInterface = NewYenAdapter(*external.NewJPService())
	yen := taxInterface.CalcTaxIncluded(dollar)

	if yen != 121 {
		t.Errorf("Error %d", yen)
	}
}
