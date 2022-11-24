package base

import (
	"bytes"
	"fmt"
	"strings"
	"testing"
)

// %q 单引号围绕的字符字面值，由Go语法安全地转义	Printf("%q", 0x4E2D)	'中'
func TestTrim(t *testing.T) {
	t.Run("TrimSpace", func(t *testing.T) {
		// TrimSpace(string)
		fmt.Println(strings.TrimSpace("\t\n\f I love Go!! \n\r")) // I love Go!!
		fmt.Println(strings.TrimSpace("I love Go!! \f\v \n\r"))   // I love Go!!
		fmt.Println(strings.TrimSpace("I love Go!!"))             // I love Go!!

		// TrimSpace([]byte)
		fmt.Printf("%q\n", bytes.TrimSpace([]byte("\t\n\f I love Go!! \n\r"))) // "I love Go!!"
		fmt.Printf("%q\n", bytes.TrimSpace([]byte("I love Go!! \f\v \n\r")))   // "I love Go!!"
		fmt.Printf("%q\n", bytes.TrimSpace([]byte("I love Go!!")))             // "I love Go!!"
	})

	fmt.Println("----------------------------------------")

	// 只修首尾的字符
	t.Run("Trim", func(t *testing.T) {
		// Trim、TrimLeft、TrimRight(string)
		fmt.Println(strings.Trim("\t\n fffI love Go!!\n \rfff", "\t\n\r f"))
		// I love Go!!
		fmt.Printf("%q\n", strings.TrimLeft("\t\n fffI love Go!!\n \rfff", "\t\n\r f"))
		// "I love Go!!\n \rfff"
		fmt.Printf("%q\n", strings.TrimRight("\t\n fffI love Go!!\n \rfff", "\t\n\r f"))
		// "\t\n fffI love Go!!"
		fmt.Println(strings.Trim("\t\n fffI ffff love Go!!\n \rfff", "\t\n\r f"))
		// I ffff love Go!!

		// Trim、TrimLeft、TrimRight([]byte)
		fmt.Printf("%q\n", bytes.Trim([]byte("\t\n fffI love Go!!\n \rfff"), "\t\n\r f"))
		// I love Go!!
		fmt.Printf("%q\n", bytes.TrimLeft([]byte("\t\n fffI love Go!!\n \rfff"), "\t\n\r f"))
		// "I love Go!!\n \rfff"
		fmt.Printf("%q\n", bytes.TrimRight([]byte("\t\n fffI love Go!!\n \rfff"), "\t\n\r f"))
		// "\t\n fffI love Go!!"
	})
}

func TestTransform(t *testing.T) {
	trans := func(r rune) rune {
		switch {
		case r == ' ':
			return ','
		}
		return r
	}
	fmt.Printf("%q\n", strings.Map(trans, "CN UK US JP"))
	fmt.Printf("%q\n", bytes.Map(trans, []byte("CN UK US JP")))
	// "CN,UK,US,JP"
	// "CN,UK,US,JP"
}
