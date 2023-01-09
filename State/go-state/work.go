package gostate

import "fmt"

type WorkStateBehavior struct {
	timer *Timer
}

func NewWorkStateBehavior(timer *Timer) *WorkStateBehavior {
	return &WorkStateBehavior{timer}
}

func (wsb *WorkStateBehavior) Start(timer *Timer) {
	timer.Second = 25 * 60
}
func (wsb *WorkStateBehavior) Stop(timer *Timer) {
	fmt.Println("stop")
}
func (wsb *WorkStateBehavior) Finish(timer *Timer) {
	timer.Second = 0
	timer.ChangeState(NewRestStateBehavior(timer))
}
