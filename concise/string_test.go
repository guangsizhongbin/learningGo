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
	rune := []rune(str2) // rune 是int32位的别称, byte是unit8的别称

	fmt.Println(reflect.TypeOf(rune[2]).Kind()) // 无论多少个字节都int32来处理
	fmt.Println(rune[2], string(rune[2]))       // 正确输出语
	fmt.Printf("%d ", len(rune))                // 其长度为4

}

// 尽管int的长度是32 bit, 但 int 与 int32 并不可以互用
func Test_intPlusint32(t *testing.T) {
	// var a int
	// var b int32
	// c := a + b
	// fmt.Println(c)
}

// Go语言中的字符串是不可变的
// 可以先转回 []byte, 再转成 string
func TestChangeChar(t *testing.T){
	// var s string = "hello"
	// s[0] = 'c' 

	var s string = "hello"
	b := []byte(s)
	s2 := string(b)
	fmt.Println(s2)
}

// 字符串的切片
func TestStringAndSlice(t *testing.T){
	s := "hello"
	s = "c" + s[1:]
	fmt.Printf("%s\n", s)
}

// 输出换行
func TestMultiLineString(t *testing.T){
	m := `hello
				world`

	fmt.Println(m)
}