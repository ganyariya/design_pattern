package gostate

import "fmt"

type RestStateBehavior struct {
	timer *Timer
}

func NewRestStateBehavior(timer *Timer) *RestStateBehavior {
	return &RestStateBehavior{timer}
}

func (wsb *RestStateBehavior) Start(timer *Timer) {
	timer.Second = 10 * 60
}
func (wsb *RestStateBehavior) Stop(timer *Timer) {
	fmt.Println("stop")
}
func (wsb *RestStateBehavior) Finish(timer *Timer) {
	timer.Second = 0
	timer.ChangeState(NewWorkStateBehavior(timer))
}
