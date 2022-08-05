package base

import (
	"bytes"
	"log"
	"os"
	"strings"
	"testing"

	"github.com/sergi/go-diff/diffmatchpatch"
)

// 文件差异对比
func TestGoDiff(t *testing.T) {
	dmp := diffmatchpatch.New()

	f1, f2 := "A 509\n", "A 509\nB 304"
	diffs := dmp.DiffMain(f1, f2, false)

	log.Println("file differences:", dmp.DiffPrettyText(diffs))

	DiffFile(f1, f2)
}

func DiffFile(before string, after string) {
	_ = os.Remove(after + ".line.diff")
	content1, _ := os.ReadFile(before)
	content2, _ := os.ReadFile(after)

	lines1 := strings.Split(string(content1), "\n")
	lines2 := strings.Split(string(content2), "\n")

	for i, line := range lines1 {
		if strings.TrimSpace(line) == "" {
			continue
		}
		dmp := diffmatchpatch.New()

		diffs := dmp.DiffMain(string(line), string(lines2[i]), false)

		if len(diffs) == 1 {
			continue
		}

		var buff bytes.Buffer
		for _, d := range diffs {
			if d.Type == diffmatchpatch.DiffInsert {
				_, _ = buff.WriteString("\x1b[32m")
				_, _ = buff.WriteString(d.Text)
				_, _ = buff.WriteString("\x1b[0m")
			} else if d.Type == diffmatchpatch.DiffDelete {
				_, _ = buff.WriteString("\x1b[31m")
				_, _ = buff.WriteString(d.Text)
				_, _ = buff.WriteString("\x1b[0m")
			} else if d.Type == diffmatchpatch.DiffEqual {
				_, _ = buff.WriteString(d.Text)
			}
		}
		AppendToFile(after+".line.diff", buff.Bytes())
	}
}

func AppendToFile(name string, data []byte) {
	file, err := os.OpenFile(name, os.O_RDWR|os.O_CREATE|os.O_APPEND, 0666)
	if err != nil {
		panic(err)
	}
	defer func() {
		_ = file.Close()
	}()

	_, err = file.Write(data)
	if err != nil {
		log.Println(err)
	}
}
