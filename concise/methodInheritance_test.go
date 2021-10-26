package concise

import (
	"fmt"
	"testing"
)

type Animal struct {
	name  string
	age   int
	phone string
}

type Monkey struct {
	Animal // 匿名字段
	school string
}

type Pig struct {
	Animal  // 匿名字段
	company string
}

// 在Human上面定义一个method
func (h *Animal) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", h.name, h.phone)
}

// Pig的method重写Animal的method
func (p *Pig) SayHi(){
    fmt.Printf("Hi I am %s, I work at %s, you can call me %s", p.name, p.company, p.phone)
}

func TestAnimal(t *testing.T) {
	feng := Monkey{Animal{"feng", 25, "123456"}, "zheda"}
	sam := Pig{Animal{"sam", 26, "123"}, "Golang Inc"}

	feng.SayHi()
	sam.SayHi()
}
