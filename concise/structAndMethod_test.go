package concise

import (
	"fmt"
	"testing"
)

type Student struct{
    name string
}

func (stu *Student) hello1 (person string) string{
    return fmt.Sprintf("hello %s, I am %s", person, stu.name)
}

func TestStruct(t *testing.T){
    stu := &Student{
        name: "Tom",
    }

    msg := stu.hello1("Jack")
    fmt.Println(msg)
}