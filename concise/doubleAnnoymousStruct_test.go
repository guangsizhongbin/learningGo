package concise

import (
	"fmt"
	"testing"
)

type HumanNew struct {
	name  string
	age   int
	phone string // Human 类型拥有的
}

type Employee struct {
	HumanNew
	speciality string
	phone      string
}

func TestStructHuman(t *testing.T) {
	Bob := Employee{HumanNew{"Bob", 32, "777-444-XXXX"}, "Designer", "333-222"}

	// 访问Bob中的phone字段
	fmt.Println("Bob's work phone is:", Bob.phone)

	// 访问Human中的phone字段
	fmt.Println("Bob's personal phone is:", Bob.HumanNew.phone)
}
