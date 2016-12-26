package value

var (
	kindToTypeFlag = map[uint]typeFlag{
		_bool:          typeBool,
		_interface:     typeInterface,
		_string:        typeString,
		_unsafePointer: typeUnsafePtr,
	}

	uintKindsToTypeFlag = map[uint]typeFlag{
		_uint:    typeUint,
		_uint8:   typeUint8,
		_uint16:  typeUint16,
		_uint32:  typeUint32,
		_uint64:  typeUint64,
		_uintptr: typeUintptr,
	}
	floatKindsToTypeFlag = map[uint]typeFlag{
		_float32: typeFloat32,
		_float64: typeFloat64,
	}
	complexKindsToTypeFlag = map[uint]typeFlag{
		_complex64:  typeComplex64,
		_complex128: typeComplex128,
	}

	intKindsToTypeFlag = map[uint]typeFlag{
		_int:   typeInt,
		_int8:  typeInt8,
		_int16: typeInt16,
		_int32: typeInt32,
		_int64: typeInt64,
	}
)
