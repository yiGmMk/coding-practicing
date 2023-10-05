package channel

import (
	"fmt"
	"testing"
	"time"
)

type user struct {
	name string
	age  int8
}

var u = user{name: "Ankur", age: 25}
var g = &u

func modifyUser(pu *user) {
	fmt.Println("modifyUser Received Vaule", pu)
	pu.name = "Anand"
}

func printUser(u <-chan *user) {
	//time.Sleep(2 * time.Second)
	fmt.Println("printUser goRoutine called", <-u)
}

// https://golang.design/go-questions/channel/principal/
// channel收发数据都是值拷贝
func TestChanVal(t *testing.T) {
	c := make(chan *user, 5)
	c <- g
	fmt.Println(g)
	// modify g
	g = &user{name: "Ankur Anand", age: 100}
	go printUser(c)
	go modifyUser(g)
	time.Sleep(1 * time.Second)
	fmt.Println(g)
}
