package sliceappend_test

import (
	"testing"

	sliceappend "github.com/ninomae42/slice_benchmark_test/slice_append"
)

func BenchmarkAppendWithNoAlloc(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sliceappend.AppendWithNoAlloc(users)
	}
}

func BenchmarkAppendWithPrealloc(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sliceappend.AppendWithPrealloc(users)
	}
}

func BenchmarkCopyWithPreallocAndIndex(b *testing.B) {
	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		sliceappend.CopyWithPreallocAndIndex(users)
	}
}
