package concise

import (
	"fmt"
	"strconv"
	"testing"
)

type Element interface{}
type List []Element

type Book struct {
	name string
	age  int
}

// 定义String方法
func (b Book) String() string {
	return "(name: " + b.name + "- age:" + strconv.Itoa(b.age) + "years)"
}


// elemet.(type)语法不能在switch外的任何逻辑里面使用,如果要在switch外面判断一个类型就使用comma-ok
func TestComma(t *testing.T) {
	list := make(List, 3)
	list[0] = 1
	list[1] = "Hello"
	list[2] = Book{"Dennis", 70}

	// if else 写法:

	// for index, element := range list {
	// 	if value, ok := element.(int); ok {
	// 		fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
	// 	} else if value, ok := element.(string); ok {
	// 		fmt.Printf("list[%d] is a string and its value is %s\n", index, value)
	// 	} else if value, ok := element.(Book); ok {
	// 		fmt.Printf("list[%d] is a Person and its value is %s\n", index, value)
	// 	} else {
	// 		fmt.Printf("list[%d] is of a different type", index)
	// 	}
	// }

	// switch重写
	for index, element := range list {
		switch value := element.(type) {
		case int:
			fmt.Printf("list[%d] is an int and its value is %d\n", index, value)
		case string:
			fmt.Printf("list[%d] is a string and its value is %s\n", index, value)
		case Person:
			fmt.Printf("list[%d] is a Person and its value is %s\n", index, value)
		default:
			fmt.Printf("list[%d] is of a different type", index)
		}

	}
}
