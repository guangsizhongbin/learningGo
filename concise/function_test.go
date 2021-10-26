package concise

import (
	"fmt"
	"testing"
)

func addTwoNum(num1 int, num2 int) int {
    return num1 + num2
}

func div(num1 int, num2 int) (int, int){
    return num1 / num2, num1 % num2
}

func TestFunction(t *testing.T){
    quo, rem := div(100, 17)
    fmt.Println(quo, rem)
    fmt.Println(addTwoNum(100, 17))
}

// 不定数量的参数
// arg ...int告诉我们GO语言这个函数接受不定数量的参数.
// 这些参数的类型全部是int, 在函数体中，变量arg是一个int的slice.
func TestVariadicFunction(t *testing.T){
    addMultiNum(1, 2, 3)
}

func addMultiNum(arg ...int){
    sum := 0
    for _, n := range arg {
        sum += n
    }

    fmt.Printf("sum is %d", sum)
}

