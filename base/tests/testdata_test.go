package tests

import (
	"bytes"
	"crypto/md5"
	"flag"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"testing"
)

var update = flag.Bool("update", false, "update .golden files")

// golden file : go test -v . -update
func TestData(t *testing.T) {
	ts := []struct {
		fileName string
		name     string
	}{
		{
			fileName: "name.golden",
			name:     "goldenfile.1",
		},
	}
	for i, v := range ts {
		h := md5.New()
		_, err := h.Write([]byte(v.name))
		if err != nil {
			t.Error(err)
		}
		hash := h.Sum(nil)

		path := filepath.Join("testdata", v.fileName)
		if *update {
			ioutil.WriteFile(path, hash, 0644)
		}
		want, err := os.ReadFile(path)
		if err != nil {
			t.Fatal(err)
		}
		if !bytes.Equal(hash, want) {
			t.Errorf("[%d]got:%s,got:%s", i, string(hash), string(want))
		}
		fmt.Printf("[%d](%s)hash:%s", i, v.name, string(hash))
	}
}
