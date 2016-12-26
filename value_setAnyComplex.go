package value

import "unsafe"

func (v *Value) setAnyComplex(tf typeFlag, dataptr unsafe.Pointer) {
	switch tf {
	case typeComplex64:
		v._complex64 = (*complex64)(dataptr)
	case typeComplex128:
		v._complex128 = (*complex128)(dataptr)
	}
}
