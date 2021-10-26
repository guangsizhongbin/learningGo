package concise

import (
	"fmt"
	"testing"  
)

type Human struct {
	name   string
	age    int
	weight int
}

type Teacher struct {
	Human      // 匿名字段，那么默认Teacher就包含了Human的所有字段
	speciality string
}

// 匿名字段可以实现字段的继承
func TestAnnymousStruct(t *testing.T) {
	// 初始化一个Teacher
	mark := Teacher{Human{"Mark", 25, 120}, "Computer Science"}

	// 访问相应的字段
	fmt.Println("His name is ", mark.name)
	fmt.Println("His age is ", mark.age)
	fmt.Println("His weight is ", mark.weight)
	fmt.Println("His speciality is ", mark.speciality)

    // 修改对应的备注信息
    mark.speciality = "AI"
    fmt.Println("Mark changed his speciality")
    fmt.Println("His speciality is", mark.speciality)

    // 修改他的年龄信息
    fmt.Println("Mark become old")
    mark.age = 46
    fmt.Println("His age is", mark.age)

    // 修改他的体重信息
    fmt.Println("Mark is not an athlet anymore")
    mark.weight += 60
    fmt.Println("His weight is", mark.weight)
}
