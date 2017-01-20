# value

test implementation of custom multi-typed storage and benchmarks vs interface{}

# Benchmarks
```
BenchmarkValueGet-8   	1000000000	         2.85 ns/op	       0 B/op	       0 allocs/op
BenchmarkEfaceGet-8   	100000000	        14.9 ns/op	       0 B/op	       0 allocs/op
BenchmarkValueSet-8   	20000000	        72.4 ns/op	       8 B/op	       1 allocs/op
BenchmarkEfaceSet-8   	30000000	        51.1 ns/op	       8 B/op	       1 allocs/op
```
