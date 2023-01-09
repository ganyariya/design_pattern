package gostate

/*
State パターンと呼ばれているが
「状態」自体は Context が持っており、 State は API (状態によって実行する処理)をもつだけ。
State という表現では「処理を実行する」ということが曖昧なので StateBehavior とここでは宣言している
*/
type TimerStateBehavior interface {
	Start(*Timer)
	Stop(*Timer)
	Finish(*Timer)
}
