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

// Panic 是一个内建函数

// recover 是一个内建函数, recover仅在延迟函数中有效，在正常的执行过程中，调用recover会返回nil,并且没有其他任何效果
// 如果当前的goroutine陷入恐慌，调用 recover 可以捕获到 panic 的输入值，并且恢复正常的执行
func TestPanicAndRecover(t *testing.T) {
	var user = os.Getenv("USER")

	// 捕获panic
	defer func(){
		if x:= recover(); x != nil{
			fmt.Print("panic handle again")
		}
	}()

	defer func(){
		if x:= recover(); x != nil{
			fmt.Println("panic handle")
			panic("There is a value for $USER")
		}
	}()


	if user == "" {
		panic("no value for $USER")
	} else {
		panic("There is a value for $USER")
	}

}
