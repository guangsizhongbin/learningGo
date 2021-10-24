package concise

import (
	"fmt"
	"testing"
)

func Test_Pointer(t *testing.T) {
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

func Test_Add(t *testing.T){
    num := 100
    add(num)
    fmt.Printf("采用传值的方式， num 的值不会增加，其值还是 %d \n", num)

    realadd(&num)
    fmt.Printf("采用传指针的方式，num 的值会增加，其值是 %d \n", num)
}
