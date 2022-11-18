package tests

import (
	"fmt"
	"testing"

	"github.com/golang/mock/gomock"
)

func TestMock(t *testing.T) {
	res, err := defaultDriveer{}.Open("mysql")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("default open mysql=", res)

	ctrl := gomock.NewController(t)

	m := NewMockDriver(ctrl)
	m.EXPECT().Open(gomock.Any()).Return("open mysql", nil).AnyTimes()

	res, err = m.Open("mysql")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("open mysql=", res)

	res, err = m.Open("redis")
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println("open redis=", res)
}
