package algorithms

import (
	"github.com/duke-git/lancet/v2/mathutil"
)

const (
	//^uint(0)        32位系统0XFFFFFFFF ; 64位系统0xFFFFFFFFFFFFFFFF
	//^uint(0) >> 63  32位为0 ;64位为1
	//32<<(^uint(0)>>63) 32位为32 ;64位为64
	_a       = uint(0)
	_b       = ^_a
	_c       = _b >> 63
	_bitSize = 32 << _c
)

// int64 8B 64bit in x64 system
type Bitmap []int64

func NewBitmap(length int64) Bitmap {
	out := make(Bitmap, length/_bitSize+1)
	return out
}

func (b *Bitmap) Set(i int64) {
	(*b)[i/_bitSize] |= 1 << uint(i%_bitSize)
}

func (b *Bitmap) Clear(i int64) {
	(*b)[i/_bitSize] &= ^(1 << uint(i%_bitSize))
}

func (b *Bitmap) ClearAll() {
	for i := range *b {
		(*b)[i] = 0
	}
}

func (b *Bitmap) Get(i int64) bool {

	return (*b)[i/_bitSize]&(1<<uint(i%_bitSize)) != 0
}

func (b *Bitmap) And(b2 Bitmap) Bitmap {
	l := mathutil.Min(len(*b), len(b2))
	out := make(Bitmap, l)
	for i := 0; i < l; i++ {
		out[i] = (*b)[i] & b2[i]
	}

	return out
}
