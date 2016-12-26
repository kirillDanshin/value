package value

import "testing"

func TestValueGet(t *testing.T) {
	v := NewValue(int64(13))
	if v == nil {
		t.Fatalf("NewValue returns nil")
	}
	if storedValue, err := v.GetAnyInt(); err != nil || storedValue != 13 {
		t.Errorf(
			"expected int64(13) and nil, got %#+v and %s. stored value type was %s",
			storedValue,
			err,
			v._type,
		)
	}
}

func BenchmarkValueGet(b *testing.B) {
	v := NewValue(int64(b.N))
	var a int64
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		a, _ = v.GetAnyInt()
	}
	_ = a
}

func BenchmarkEfaceGet(b *testing.B) {
	v := interface{}(int64(b.N))
	var a int64
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		a = v.(int64)
	}
	_ = a
}

func BenchmarkValueSet(b *testing.B) {
	a := int64(b.N)
	var v = &Value{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v.Set(a)
	}
	_ = v
}

func BenchmarkEfaceSet(b *testing.B) {
	a := int64(b.N)
	var v interface{}
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		v = interface{}(a)
	}
	_ = v
}
