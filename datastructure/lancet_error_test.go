package datastructure

import (
	"log"
	"strconv"
	"testing"

	"github.com/duke-git/lancet/v2/xerror"
	"github.com/go-playground/assert/v2"
)

func TestLancetError(t *testing.T) {
	_, err := strconv.Atoi("4o2")
	log.Printf("1.%+v", err)
	defer func() {
		v := recover()
		assert.Equal(t, err.Error(), v.(*strconv.NumError).Error())
		log.Printf("2.%+v\n", v)
	}()

	xerror.Unwrap(strconv.Atoi("4o2"))
}
