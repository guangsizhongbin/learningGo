package concise

import (
	"fmt"
	"testing"
)


// Go语言中 string, slice, map这三种类型的的实现机制类似指针，所有可以直接传递，而不有地址后传递指针
// 若函数需要改变 slice 的长度，　则仍需要取地址传递指针
// slice 改变长度时，可能返回新的slice(当引用空间不足时), 虽然是传递的引用
func TestPointer(t *testing.T) {
	str := "Golang"
    var p *string = &str // p是指向str的指针
    *p = "Hello"
    fmt.Println(str)     // Hello 修改了 p, str的值　也发生了改变
}

func add(num int){
    num += 1
}

func realadd(num *int){
    *num += 1
}

func TestAdd(t *testing.T){
    num := 100
    add(num)
    fmt.Printf("采用传值的方式， num 的值不会增加，其值还是 %d \n", num)

    realadd(&num)
    fmt.Printf("采用传指针的方式，num 的值会增加，其值是 %d \n", num)
}
