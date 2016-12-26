package value

// Get value of any type
func (v *Value) Get() interface{} {
	switch v._type {
	case typeByte:
		return *v._byte
	case typeBytePtr:
		return v._byte
	case typeComplex128:
		return *v._complex128
	case typeComplex128Ptr:
		return v._complex128
	case typeComplex64:
		return *v._complex64
	case typeComplex64Ptr:
		return v._complex64
	case typeError:
		return *v._error
	case typeFloat32:
		return *v._float32
	case typeFloat32Ptr:
		return v._float32
	case typeFloat64:
		return *v._float64
	case typeFloat64Ptr:
		return v._float64
	case typeInt:
		return *v._int
	case typeIntPtr:
		return v._int
	case typeInt16:
		return *v._int16
	case typeInt16Ptr:
		return v._int16
	case typeInt32:
		return *v._int32
	case typeInt32Ptr:
		return v._int32
	case typeInt64:
		return *v._int64
	case typeInt64Ptr:
		return v._int64
	case typeInt8:
		return *v._int8
	case typeInt8Ptr:
		return v._int8
	// case typeRune:
	// 	return v._rune
	// case typeRunePtr:
	// 	return v._rune
	case typeString:
		return *v._string
	case typeStringPtr:
		return v._string
	case typeUint:
		return *v._uint
	case typeUintPtr:
		return v._uint
	case typeUint16:
		return *v._uint16
	case typeUint16Ptr:
		return v._uint16
	case typeUint32:
		return *v._uint32
	case typeUint32Ptr:
		return v._uint32
	case typeUint64:
		return *v._uint64
	case typeUint64Ptr:
		return v._uint64
	case typeUintptr:
		return *v._uintptr
	case typeUintptrPtr:
		return v._uintptr

	default:
		return v._eface
	}
}
