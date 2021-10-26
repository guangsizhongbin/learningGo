package concise

import (
	"fmt"
	"testing"
)

// init函数(能够应用于所有的package), init函数是可选的
// main函数(只能应用于package main), package main就必须包含一个main函数
// 这两个函数在定义时不能有任何的参数和返回值
func TestMain(t *testing.T) {
	fmt.Println("Hello TestMain")
}

func TestInit(t *testing.T) {
	fmt.Println("Hello TestInit")
}
