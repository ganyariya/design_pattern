package gostate

/*
いわゆる Context
複数の状態を持つ Entity クラス

内部で State を持ち、それぞれの State に処理を移譲する

「状態」（たとえばタイマーの秒数など)は context 側で持つ。
State Interface らは「秒数」などを持たず、 context メンバを受け取る。
よって、 Work/Rest の State らは処理を実行するのみであり、状態自体は Context (Timer) が持つことに注意。
*/
type Timer struct {
	StateBehavior TimerStateBehavior
	Second        int
}

func NewTimer() *Timer {
	return &Timer{Second: 0}
}
func (t *Timer) ChangeState(stateBehavior TimerStateBehavior) {
	t.StateBehavior = stateBehavior
}
func (t *Timer) Start() {
	t.StateBehavior.Start(t)
}
func (t *Timer) Stop() {
	t.StateBehavior.Stop(t)
}
func (t *Timer) Finish() {
	t.StateBehavior.Finish(t)
}
