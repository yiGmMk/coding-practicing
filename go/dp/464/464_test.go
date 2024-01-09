package jzoffer

import "testing"

func Test464(t *testing.T) {
	if got := canIWin(10, 11); got != false {
		t.Errorf("10, 11 got:%v, want:%v", got, false)
	}

	if got := canIWin(10, 40); got != false {
		t.Errorf("10, 41 got:%v, want:%v", got, false)
	}

	if got := canIWin(10, 10); got != true {
		t.Errorf("10, 10 got:%v, want:%v", got, true)
	}
}
