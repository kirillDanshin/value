package value

import "unsafe"

func (v *Value) setAnyFloat(tf typeFlag, dataptr unsafe.Pointer) {
	switch tf {
	case typeFloat32:
		v._float32 = (*float32)(dataptr)
	case typeFloat64:
		v._float64 = (*float64)(dataptr)
	}
}
