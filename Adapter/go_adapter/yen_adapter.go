package go_adapter

import "github.com/ganyariya/design_pattern/Adapter/go_adapter/external"

type YenAdapter struct {
	jpService external.JPService
}

func NewYenAdapter(jp external.JPService) *YenAdapter {
	return &YenAdapter{jp}
}

func (ya *YenAdapter) CalcTaxIncluded(dollar int) int {
	yen := dollar * 110
	return ya.jpService.TaxIncluded(yen)
}
