package concise

import (
	"fmt"
	"testing"
)

type Fruit struct {
	name  string
	age   int
	phone string
}

type Apple struct {
	Fruit
	school string
	loan   float32
}

type Banana struct {
	Fruit
	company string
	money   float32
}

func (f Fruit) SayHi() {
	fmt.Printf("Hi, I am %s you can call me on %s\n", f.name, f.phone)
}

func (h Fruit) Sing(lyrics string) {
	fmt.Println("La la la la la ...", lyrics)
}

// Apple 重载SayHi
func (a Apple) SayHi() {
	fmt.Printf("Hi, I am %s, I work at %s. Call me on %s \n", a.name, a.school, a.phone)
}

// Banana 重载SayHi
func (b Banana) SayHi() {
	fmt.Println("Hi, I am Banana")
}

type Men interface {
	SayHi()
	Sing(lyrics string)
}

// interface 就是一组抽象方法的集合，必须由其他非interface类型实现，而不能自我实现，Go语言interface实现了duck-typing

func TestFruitSayHi(t *testing.T) {
    mike := Banana{Fruit{"mike", 12, "12345"}, "tx", 123}
    feng := Banana{Fruit{"feng", 12, "12345"}, "tx", 123}

    liu := Apple{Fruit{"liu", 23, "678990"},"hf",  32}
    zhuzhu := Apple{Fruit{"zhuzhu", 23, "678990"},"zheda",  32}

    // 定义Men类型的变量i
    var i Men

    // i 可以存储mike
    i = mike
    fmt.Println("This is Mike, a banana:")
    i.SayHi()
    i.Sing("November rain")

    // i　也可以存储Banana
    i = liu
    fmt.Println("This is liu, An apple:")
    i.SayHi()
    i.Sing("Born to be wild")

    // 定义slice Men
    fmt.Println("Let's use a slice of Men and see what happens")
    x := make([]Men , 4)

    x[0], x[1], x[2], x[3] = mike, feng, liu, zhuzhu
    for _, value := range x{
        value.SayHi()
    }
}
