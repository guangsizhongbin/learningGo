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