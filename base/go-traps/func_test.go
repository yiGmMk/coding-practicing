package base

import (
	"fmt"
	"testing"
)

type NameSeter interface {
	SetName(name string)
}

type CusType struct {
	Name string
}

func (c CusType) SetName(name string) {
	// c.Name = name // ineffective assignment to field CusType.Name
}

func (c *CusType) SetNameWithPointer(name string) {
	c.Name = name
}

type CusType2 struct {
	Name string
}

func (c *CusType2) SetName(name string) {
	c.Name = name
}

func TestFunc(t *testing.T) {
	var c NameSeter
	c = CusType{}
	c.SetName("aaa")
	fmt.Println(c)

	// c = CusType2{} //cannot use (CusType2 literal) (value of type CusType2) as NameSeter value in assignment: CusType2 does not implement
	c = &CusType2{}
	c.SetName("bbb")
	fmt.Println(c)

	c = &CusType{}
	c.SetName("ccc")
	fmt.Println(c)
	// {}
	// &{bbb}
	// &{}
}
