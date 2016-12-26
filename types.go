package value

import (
	"unsafe"
)

type (
	typeFlag int

	// Value holds pointers to value of many types.
	// if Value type is one of built-in types,
	// it will be stored right in *<type> field.
	// otherwise it will be stored in *interface{}
	// field that makes lower performance.
	Value struct {
		_type       typeFlag // should be first for faster access
		_bool       *bool
		_byte       *byte
		_complex128 *complex128
		_complex64  *complex64
		_error      *error
		_float32    *float32
		_float64    *float64
		_int        *int
		_int16      *int16
		_int32      *int32
		_int64      *int64
		_int8       *int8
		_rune       *rune
		_string     *string
		_uint       *uint
		_uint16     *uint16
		_uint32     *uint32
		_uint64     *uint64
		_uint8      *uint8
		_uintptr    *uintptr
		_unsafePtr  *unsafe.Pointer
		_eface      interface{}
	}

)
