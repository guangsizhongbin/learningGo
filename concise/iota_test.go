package concise

import (
	"fmt"
	"testing"
)

// iota默认开始值是0, 每调用一次加1
// 每遇到一个const关键字，iota就会重置
func TestIota(t *testing.T) {
	const (
		x = iota
		y = iota
		z = iota
		w
	)

	const v = iota

	fmt.Printf("x, y, z, w, v的值分别是%d, %d, %d, %d, %d", x, y, z, w, v)
}
