# Go slice benchmark test

## Description

This is a benchmark test for slice copy implementation with append or index access copy.

## Command

```bash
make test
```

## Result

```txt
❯ make test
Running tests...
go test -benchmem -bench=. ./...
goos: darwin
goarch: arm64
pkg: github.com/ninomae42/slice_benchmark_test/slice_append
BenchmarkAppendWithNoAlloc-10                     250782              5098 ns/op           50416 B/op         12 allocs/op
BenchmarkAppendWithPrealloc-10                    673382              1804 ns/op           16384 B/op          1 allocs/op
BenchmarkCopyWithPreallocAndIndex-10              715849              1702 ns/op           16384 B/op          1 allocs/op
PASS
ok      github.com/ninomae42/slice_benchmark_test/slice_append  4.901s
```
