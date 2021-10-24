package concise

import (
	"fmt"
	"reflect"
	"testing"
)

func Test_Variable(t *testing.T) {
	var a, b, c = 1, 2, 3
	e, f, g := 4, 5, 6
	fmt.Println(a, b, c)
	fmt.Println(e, f, g)
}

func Test_String(t *testing.T) {
	str1 := "golang"
	str2 := "Go语言"

	fmt.Println(reflect.TypeOf(str2[2]).Kind()) // 字符串是以byte数组类型保存，类型是unit8,占一个字节
	fmt.Println(str1[2], string(str1[2]))       // 打印时需要用string(), 否则打印的是编码值
	fmt.Printf("%d %c\n", str2[2], str2[2])     // str2[2] 不是语
	fmt.Println("len(str2)", len(str2))         // len(str2) 是8，而不是4，GO占2字节，语言占6字节
}

func Test_String_rune(t *testing.T) {
	str2 := "Go语言"
	rune := []rune(str2)

	fmt.Println(reflect.TypeOf(rune[2]).Kind()) // 无论多少个字节都int32来处理
	fmt.Println(rune[2], string(rune[2]))       // 正确输出语
	fmt.Printf("%d ", len(rune))                // 其长度为4

}
