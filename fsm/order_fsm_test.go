package fsm

import (
	"log"
	"testing"

	"github.com/looplab/fsm"
)

var (
	_emptyOrderEvents = fsm.Events{
		{Name: "book", Src: []string{"empty"}, Dst: "booking"},
		{Name: "pay", Src: []string{"booking"}, Dst: "paid"},
		{Name: "cancel", Src: []string{"booking", "paid"}, Dst: "cancelled"},
	}

	_emptyOrderEvents1 = []fsm.EventDesc{
		{Name: "book", Src: []string{""}, Dst: "booking"},
		{Name: "pay", Src: []string{"booking"}, Dst: "paid"},
		{Name: "cancel", Src: []string{"booking", "paid"}, Dst: "cancelled"},
	}
)

func TestOrder(t *testing.T) {
	fsm := fsm.NewFSM(
		"empty",
		_emptyOrderEvents,
		fsm.Callbacks{
			"leave_state": func(e *fsm.Event) { log.Printf("\nleave state:%s,cur state:%s\n", e.Src, e.Dst) },
		})

	log.Println(fsm.AvailableTransitions())

	err := fsm.Event("book")
	if err != nil {
		log.Println(err)
	}

	err = fsm.Event("pay")
	if err != nil {
		log.Println(err)
	}

	err = fsm.Event("cancel")
	if err != nil {
		log.Println(err)
	}
}

func TestOrderEmptyInit(t *testing.T) {
	// 初始状态为空时,一个transaction会进两次callback
	fsm := fsm.NewFSM(
		"",
		_emptyOrderEvents1,
		fsm.Callbacks{
			"leave_state": func(e *fsm.Event) { log.Printf("\nleave state:%s,cur state:%s\n", e.Src, e.Dst) },
		})

	log.Println(fsm.AvailableTransitions())

	err := fsm.Event("book")
	if err != nil {
		log.Println(err)
	}

	err = fsm.Event("pay")
	if err != nil {
		log.Println(err)
	}

	err = fsm.Event("cancel")
	if err != nil {
		log.Println(err)
	}
}
