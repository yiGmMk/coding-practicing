package designpattern

import (
	"fmt"
	"testing"
)

func TestNewGun(t *testing.T) {
	ak47, _ := NewGun("ak47")
	musket, _ := NewGun("musket")

	if _, ok := ak47.(IGun); !ok {
		t.Errorf("wrong type")
	}

	printDetails(ak47)
	printDetails(musket)
}

func printDetails(g IGun) {
	fmt.Printf("Gun: %s", g.getName())
	fmt.Println()
	fmt.Printf("Power: %d", g.getPower())
	fmt.Println()
}
