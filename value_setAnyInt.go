package value

import "unsafe"

func (v *Value) setAnyInt(tf typeFlag, dataptr unsafe.Pointer) {
	switch tf {
	case typeInt:
		v._int = (*int)(dataptr)
	case typeInt8:
		v._int8 = (*int8)(dataptr)
	case typeInt16:
		v._int16 = (*int16)(dataptr)
	case typeInt32:
		v._int32 = (*int32)(dataptr)
	case typeInt64:
		v._int64 = (*int64)(dataptr)
	}
}
