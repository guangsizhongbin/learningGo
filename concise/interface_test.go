package concise

import (
	"fmt"
	"testing"
)

type Person interface {
    getName() string
}

type Studentnew struct{
    name string
    age int
}

func (stu *Studentnew) getName() string{
    return stu.name
}

type Worker struct {
    name string
    gender string
}

func (w *Worker) getName() string{
    return w.name
}

func Test_Interface(t *testing.T) {
    var p Person = &Studentnew{
        name: "Tom",
        age : 18,
    }

    var w Person = &Worker{
        name: "feng",
        gender: "unknow",
    }

    fmt.Println(p.getName())
    fmt.Println(w.getName())
}

// 空interface(interface{}) 不包含任何的method
// 正因为如此，所有的类型都实现了空interface
// 空interface对于描述起不到任何的作用
func Test_EmptyInterface(t *testing.T){
    m := make(map[string]interface{})
    m["name"] = "Tom"
    m["age"] = 18
    m["scores"] = [3]int{98, 99, 85}
    fmt.Println(m)
}