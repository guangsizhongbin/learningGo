package concise

import (
	"fmt"
	"testing"
)

type Student struct {
	name string
}

func (stu *Student) hello1(person string) string {
	return fmt.Sprintf("hello %s, I am %s", person, stu.name)
}

func TestStruct(t *testing.T) {
	stu := &Student{
		name: "Tom",
	}

	msg := stu.hello1("Jack")
	fmt.Println(msg)
}

type person struct {
	name string
	age  int
}

func TestInitStudent(t *testing.T) {
	var P person // p现在就是person类型
    P.name = "Astaxie"
    P.age = 25
    fmt.Printf("The person's name is %s and age is %d \n", P.name, P.age)

    P1 := person{"Tom", 25} // 按照顺序提供初始化值
    fmt.Printf("The P1's name is %s and age is %d \n", P1.name, P1.age)

    P2 := person{age:24, name:"Tom"}
    fmt.Printf("The P2's name is %s and age is %d \n", P2.name, P2.age)
}
