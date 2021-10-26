package concise

import (
	"fmt"
	"testing"
)

type Gender int8

const (
	MALE   Gender = 1
	FEMALE Gender = 2
)

func Test_switch(t *testing.T) {

	gender := MALE
	switchWithBreak(gender)

	switchfallthrough(gender)

}

func switchWithBreak(gender Gender) {
	switch gender {
	case FEMALE:
		fmt.Println("female")
	case MALE:
		fmt.Println("male")
	default:
		fmt.Println("unknown")
	}
}

func switchfallthrough(gender Gender) {
	switch gender {
	case FEMALE:
		fmt.Println("female")
		fallthrough
	case MALE:
		fmt.Println("female")
		fallthrough
	default:
		fmt.Println("unknown")
	}
}

// Go 语言里面switch默认相当于每个case最后带brek, 匹配成功后不会自动向下执行其他case
// 而是跳出整个switch, 但是可以使用fallthrough强制执行后面的case代码
func TestSwitchWithFallthrough(t *testing.T) {
	i := 10
	switch i {
	case 1:
		fmt.Println("i is equal to 1")
		fallthrough
	case 2, 3, 4:
		fmt.Println("i is equal to 2, 3, or 4")
		fallthrough
	case 10:
		fmt.Println("i is equal to 10")
		fallthrough
	default:
		fmt.Println("All I know is that i is an interger")
	}
}
