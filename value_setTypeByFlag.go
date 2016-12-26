package value

import "unsafe"

func (v *Value) setByTypeFlag(t typeFlag, value interface{}) {
	if t == typeInterface {
		v._eface = value
		return
	}
	switch t {
	case typeByte, typeBytePtr:
		v._byte = (*byte)(data(value))
		return
	case typeComplex128, typeComplex128Ptr:
		v._complex128 = (*complex128)(data(value))
		return
	case typeComplex64, typeComplex64Ptr:
		v._complex64 = (*complex64)(data(value))
		return
	case typeError:
		v._error = (*error)(data(value))
		return
	case typeFloat32, typeFloat32Ptr:
		v._float32 = (*float32)(data(value))
		return
	case typeFloat64, typeFloat64Ptr:
		v._float64 = (*float64)(data(value))
		return
	case typeInt, typeIntPtr:
		v._int = (*int)(data(value))
		return
	case typeInt16, typeInt16Ptr:
		v._int16 = (*int16)(data(value))
		return
	case typeInt32, typeInt32Ptr:
		v._int32 = (*int32)(data(value))
		return
	case typeInt64, typeInt64Ptr:
		v._int64 = (*int64)(data(value))
		return
	case typeInt8, typeInt8Ptr:
		v._int8 = (*int8)(data(value))
		return
	case typeRune, typeRunePtr:
		v._rune = (*rune)(data(value))
		return
	case typeString, typeStringPtr:
		v._string = (*string)(data(value))
		return
	case typeUint, typeUintPtr:
		v._uint = (*uint)(data(value))
		return
	case typeUint16, typeUint16Ptr:
		v._uint16 = (*uint16)(data(value))
		return
	case typeUint32, typeUint32Ptr:
		v._uint32 = (*uint32)(data(value))
		return
	case typeUint64, typeUint64Ptr:
		v._uint64 = (*uint64)(data(value))
		return
	case typeUint8, typeUint8Ptr:
		v._uint8 = (*uint8)(data(value))
		return
	case typeUintptr, typeUintptrPtr:
		v._uintptr = (*uintptr)(data(value))
		return
	case typeUnsafePtr, typeUnsafePtrPtr:
		v._unsafePtr = (*unsafe.Pointer)(data(value))
		return
	}
}
