package value

// NewValue instantiates a new Value for storage instance
func NewValue(yourValue interface{}) *Value {
	v := &Value{}
	v.Set(yourValue)

	return v
}
