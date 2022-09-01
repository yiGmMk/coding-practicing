package base

import (
	"fmt"
	"testing"
	"time"
)

func TestFormat(t *testing.T) {
	v := fmt.Sprintln(time.Now().Format("2006年01月02日 15时04分05秒"))
	fmt.Println(v)

	// 2022年09月01日 09时50分38秒
}
