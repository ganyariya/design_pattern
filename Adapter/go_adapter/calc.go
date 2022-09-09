package go_adapter

// 開発アプリ側で用意しているインターフェース
// このインターフェースに合うように 外部ライブラリをラップする Adapter を用意する
type TaxInterface interface {
	CalcTaxIncluded(int) int
}
