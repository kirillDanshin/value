package value

import "unsafe"

func kind(eface interface{}) uint {
	return uint((*emptyInterface)(unsafe.Pointer(&eface)).typ.kind & kindMask)
}

func data(eface interface{}) unsafe.Pointer {
	// return unsafe.Pointer((*[2]unsafe.Pointer)(unsafe.Pointer(&eface))[1])
	return (*emptyInterface)(unsafe.Pointer(&eface)).word
}
