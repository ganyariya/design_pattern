package external

// 外部ライブラリであるため変更できない
// よって この Service を Adapter でラップする

// 日本円の税込価格額を計算する外部ライブラリ

const taxRate = 0.1

type JPService struct {
}

func NewJPService() *JPService {
	return &JPService{}
}

func (*JPService) TaxIncluded(yen int) int {
	return int(taxRate*float32(yen)) + yen
}
