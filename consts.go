package value

const kindMask = (1 << 5) - 1

const (
	typeByte typeFlag = iota
	typeBool
	typeComplex128
	typeComplex64
	typeError
	typeFloat32
	typeFloat64
	typeInt
	typeInt16
	typeInt32
	typeInt64
	typeInt8
	typeRune
	typeString
	typeUint
	typeUint16
	typeUint32
	typeUint64
	typeUint8
	typeUintptr
	typeUnsafePtr
	typeInterface

	typeBytePtr
	typeComplex128Ptr
	typeComplex64Ptr
	typeFloat32Ptr
	typeFloat64Ptr
	typeIntPtr
	typeInt16Ptr
	typeInt32Ptr
	typeInt64Ptr
	typeInt8Ptr
	typeRunePtr
	typeStringPtr
	typeUintPtr
	typeUint16Ptr
	typeUint32Ptr
	typeUint64Ptr
	typeUint8Ptr
	typeUintptrPtr
	typeUnsafePtrPtr
)

func (t typeFlag) String() string {
	switch t {
	case typeByte:
		return "typeByte"
	case typeBool:
		return "typeBool"
	case typeComplex128:
		return "typeComplex128"
	case typeComplex64:
		return "typeComplex64"
	case typeError:
		return "typeError"
	case typeFloat32:
		return "typeFloat32"
	case typeFloat64:
		return "typeFloat64"
	case typeInt:
		return "typeInt"
	case typeInt16:
		return "typeInt16"
	case typeInt32:
		return "typeInt32"
	case typeInt64:
		return "typeInt64"
	case typeInt8:
		return "typeInt8"
	case typeRune:
		return "typeRune"
	case typeString:
		return "typeString"
	case typeUint:
		return "typeUint"
	case typeUint16:
		return "typeUint16"
	case typeUint32:
		return "typeUint32"
	case typeUint64:
		return "typeUint64"
	case typeUint8:
		return "typeUint8"
	case typeUintptr:
		return "typeUintptr"
	case typeUnsafePtr:
		return "typeUnsafePtr"
	case typeInterface:
		return "typeInterface"

	case typeBytePtr:
		return "typeBytePtr"
	case typeComplex128Ptr:
		return "typeComplex128Ptr"
	case typeComplex64Ptr:
		return "typeComplex64Ptr"
	case typeFloat32Ptr:
		return "typeFloat32Ptr"
	case typeFloat64Ptr:
		return "typeFloat64Ptr"
	case typeIntPtr:
		return "typeIntPtr"
	case typeInt16Ptr:
		return "typeInt16Ptr"
	case typeInt32Ptr:
		return "typeInt32Ptr"
	case typeInt64Ptr:
		return "typeInt64Ptr"
	case typeInt8Ptr:
		return "typeInt8Ptr"
	case typeRunePtr:
		return "typeRunePtr"
	case typeStringPtr:
		return "typeStringPtr"
	case typeUintPtr:
		return "typeUintPtr"
	case typeUint16Ptr:
		return "typeUint16Ptr"
	case typeUint32Ptr:
		return "typeUint32Ptr"
	case typeUint64Ptr:
		return "typeUint64Ptr"
	case typeUint8Ptr:
		return "typeUint8Ptr"
	case typeUintptrPtr:
		return "typeUintptrPtr"
	case typeUnsafePtrPtr:
		return "typeUnsafePtrPtr"
	}
	return "unknown"
}
