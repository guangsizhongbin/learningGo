package concise

import (
	"fmt"
	"testing"
)

// 1. 通过fmt输出interface中存储的类型
func TestFmtPrintInterface(t *testing.T) {
	v := "hello world"
	fmt.Println(typeof(v))
}

// func typeof(v interface{}) string {
// 	return fmt.Sprintf("%T", v)
// }

// 2. 通过反射获取interface中存储的类型
func TestRefectInterface(t *testing.T) {
	v := "hello world"
	fmt.Println(typeof(v))
}

// func typeof(v interface{}) string {
// 	return reflect.TypeOf(v).String()
// }

// 3. 通过断言判断
// `value, ok = element.(T)`, 这里 `value` 就是变量的值, `ok` 是一个 `bool` 类型, `element` 是 `interface` 变量, `T`
func TestAssertInterface(t *testing.T) {
	v := "hello world"
	fmt.Println(typeof(v))
}

func typeof(v interface{}) string {
	switch t := v.(type) {
	case int:
		return "int"
	case float64:
		return "float64"
	case string:
		return "string"
	default:
		_ = t
		return "unknown"
	}
}
