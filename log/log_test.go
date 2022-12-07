package log

import (
	"fmt"
	"os"
	"os/exec"
	"testing"
)

func TestPath(t *testing.T) {
	t.Run("GetWd", func(t *testing.T) {
		dir, err := os.Getwd()
		if err != nil {
			t.Errorf("GetWd:%v", err)
		}
		fmt.Println("GetWd", dir)
	})

	t.Run("os.Args[0]", func(t *testing.T) {
		dir, err := exec.LookPath(os.Args[0])
		if err != nil {
			t.Errorf("os.Args:%v", err)
		}
		fmt.Println(dir)
	})

	t.Run("runtime.Caller", func(t *testing.T) {

	})
}
