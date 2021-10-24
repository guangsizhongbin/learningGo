package concise

import (
	"errors"
	"fmt"
	"os"
	"testing"
)

// error 是指能够预知到的错误
func TestErrorHandle(t *testing.T) {
	_, err := os.Open("filename.txt")
	if err != nil {
		fmt.Println(err)
	}
}

// 自定义错误
func hello(name string) error {
	if len(name) == 0 {
		return errors.New("error: name is null")
	}
	fmt.Println("Hello,", name)
	return nil
}

func TestHelloError(t *testing.T) {
	if err := hello(""); err != nil {
		fmt.Println(err)
	}
}

// panic是不可预知到的错误

func get(index int) (ret int) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Some error happened!", r)
			ret = -1
		}
	}()

	arr := [3]int{2, 3, 4}
	return arr[index]
}

func TestPanic(t *testing.T) {
	fmt.Println(get(5))
	fmt.Println("finished")
}
