package value

import "unsafe"

func (v *Value) setAnyUint(tf typeFlag, dataptr unsafe.Pointer) {
	switch tf {
	case typeUint:
		v._uint = (*uint)(dataptr)
	case typeUint8:
		v._uint8 = (*uint8)(dataptr)
	case typeUint16:
		v._uint16 = (*uint16)(dataptr)
	case typeUint32:
		v._uint32 = (*uint32)(dataptr)
	case typeUint64:
		v._uint64 = (*uint64)(dataptr)
	case typeUintptr:
		v._uintptr = (*uintptr)(dataptr)
	}
}
