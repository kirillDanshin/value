package value

// GetAnyInt returns any int* value as int64
// this method will not return uint* values
// because of overflow danger
func (v *Value) GetAnyInt() (int64, error) {
	switch v._type {
	case typeInt, typeIntPtr:
		return int64(*v._int), nil
	case typeInt8, typeInt8Ptr:
		return int64(*v._int8), nil
	case typeInt16, typeInt16Ptr:
		return int64(*v._int16), nil
	case typeInt32, typeInt32Ptr:
		return int64(*v._int32), nil
	case typeInt64, typeInt64Ptr:
		return *v._int64, nil
	}
	return 0, errNotAnInt
}
