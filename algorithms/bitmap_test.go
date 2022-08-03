package algorithms

import (
	"log"
	"math/rand"
	"testing"
	"unicode/utf8"
	"unsafe"

	datastructure "github.com/duke-git/lancet/v2/datastructure/set"
)

func TestBitmap1(t *testing.T) {
	max := (1_0000_0000)
	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 100, 1000, 1000_000, max}
	b := NewBitmap(max)
	for _, v := range nums {
		b.Set(v)
	}

	log.Println(
		"len:", len(b),
		"size of b:", unsafe.Sizeof(b[:]),
		"cap:", cap(b), "Bytes:", cap(b)*8, "MB:", float64(cap(b)*8)/1024/1024,
		"origin MB:", max*8/1024/1024.0)

	for _, v := range nums {
		if !b.Get(v) {
			t.Errorf("%d not in set", v)
		}
	}
}

func TestBitmap2(t *testing.T) {
	max := (1_0000_0000)
	notIn := datastructure.NewSet[int]()
	b := NewBitmap(max)
	for i := (0); i < max; i++ {
		if rand.Int()%555 == 3 {
			notIn.Add(i)
			continue
		}
		b.Set(i)
	}

	log.Println("len:", len(b), "size of b:", unsafe.Sizeof(b))

	for i := (0); i < max; i++ {
		if !b.Get(i) && !notIn.Contain(i) {
			t.Errorf("%d should be member of Bitmap", i)
		}
	}
	log.Println("ok")
}

func TestBitMapFromregexp(t *testing.T) {
	// copy from regexp/regexp.go
	// Bitmap used by func special to check whether a character needs to be escaped.
	var specialBytes [16]byte

	// special reports whether byte b needs to be escaped by QuoteMeta.
	special := func(b byte) bool {
		return b < utf8.RuneSelf && specialBytes[b%16]&(1<<(b/16)) != 0
	}

	for _, b := range []byte(`\.+*?()|[]{}^$`) {
		specialBytes[b%16] |= 1 << (b / 16)
	}

	for _, b := range []byte(`abcd,./edf123\.+*?()|[]{}^$`) {
		log.Println(string(b), ":", special(b))
	}
}

func TestBitmap3(t *testing.T) {

	max := (1_0000_0000)
	notIn := datastructure.NewSet[int]()
	b := NewBitmap(max)

	// 另一种实现,在标准库中的regexp/regexp.go中
	var specialBytes [1_0000_0000/8 + 1]byte
	specialBytesLen := 1_0000_0000/8 + 1
	special := func(b int) bool {
		return specialBytes[b%specialBytesLen]&(1<<(b/specialBytesLen)) != 0
	}
	setSpecial := func(b int) {
		specialBytes[b%specialBytesLen] |= 1 << (b / specialBytesLen)
	}

	for i := (0); i < max; i++ {
		if rand.Int()%555 == 3 {
			notIn.Add(i)
			continue
		}
		b.Set(i)
		setSpecial(i)
	}

	log.Println("len:", len(b), "size of b:", unsafe.Sizeof(b))

	for i := (0); i < max; i++ {
		if !b.Get(i) && !notIn.Contain(i) {
			t.Errorf("%d should be member of Bitmap", i)
		}
		if !special(i) && !notIn.Contain(i) {
			t.Errorf("%d should be member of specialBytes", i)
		}
	}
	log.Println("ok")
}
