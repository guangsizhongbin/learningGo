package test

import "testing"

func Fib(n int) int {
	if n < 2 {
		return n
	}
	return Fib(n-1) + Fib(n-2)
}

// 简单的test
// 用 go test . 进行测试
func TestFib(t *testing.T) {
	var (
		in       = 7
		expected = 13
	)

	actual := Fib(in)
	if actual != expected {
		t.Errorf("Fib(%d) = %d; expected %d", in, actual, expected)
	}
}
