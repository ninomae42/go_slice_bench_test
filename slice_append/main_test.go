package sliceappend_test

import (
	"strconv"
	"testing"

	sliceappend "github.com/ninomae42/slice_benchmark_test/slice_append"
)

const testSize = 1000

var users = []sliceappend.User{}

func TestMain(m *testing.M) {
	for i := 0; i < testSize; i++ {
		user := sliceappend.User{ID: "user" + strconv.Itoa(i)}
		users = append(users, user)
	}

	m.Run()
}
