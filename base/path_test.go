package base

import (
	"log"
	"path/filepath"
	"testing"
)

func TestPath(t *testing.T) {
	strs := []string{"./go-job"}
	for _, item := range strs {
		dirAbs, err := filepath.Abs(item)
		if err != nil {
			log.Println(err)
			return
		}
		pkg := filepath.Base(dirAbs)
		log.Println(pkg)
		if pkg != item[2:] {
			t.Errorf("pkg:%s not right", pkg)
		}
	}
}
