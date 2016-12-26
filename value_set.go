package value

// Set a value
func (v *Value) Set(yourValue interface{}) {
	if tf, ok := intKindsToTypeFlag[kind(yourValue)]; ok {
		v._type = tf
		v.setAnyInt(tf, data(yourValue))
		return
	}
	if tf, ok := complexKindsToTypeFlag[kind(yourValue)]; ok {
		v._type = tf
		v.setAnyComplex(tf, data(yourValue))
		return
	}
	if tf, ok := floatKindsToTypeFlag[kind(yourValue)]; ok {
		v._type = tf
		v.setAnyFloat(tf, data(yourValue))
		return
	}
	if tf, ok := uintKindsToTypeFlag[kind(yourValue)]; ok {
		v._type = tf
		v.setAnyUint(tf, data(yourValue))
		return
	}
	if tf, ok := kindToTypeFlag[kind(yourValue)]; ok {
		v._type = tf
		v.setAnyOther(tf, data(yourValue))
		return
	}
	v._type = typeInterface
	v._eface = yourValue
}
