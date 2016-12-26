package value

import "unsafe"

func (v *Value) setAnyOther(tf typeFlag, dataptr unsafe.Pointer) {
	switch tf {
	case typeBool:
		v._bool = (*bool)(dataptr)
	case typeInterface:
		v._eface = (*interface{})(dataptr)
	case typeString:
		v._string = (*string)(dataptr)
	case typeUnsafePtr:
		v._unsafePtr = (*unsafe.Pointer)(dataptr)
	}
}
