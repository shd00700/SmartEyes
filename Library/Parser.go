package Library

import (
	"fmt"
	"unsafe"
)

func Parser(c []uint16) {
	v := [2]uint16{c[34], c[35]}
	a := unsafe.Pointer(&v)
	b := *(*float32)(a)
}
