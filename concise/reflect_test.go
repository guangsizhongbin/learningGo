package concise

import (
	"fmt"
	"reflect"
	"testing"
)

func TestReflect(t *testing.T) {
	var x float64 = 3.4
	v := reflect.ValueOf(x)

	// 获取其类型定义
	fmt.Println("type: ", v.Type())

	// 获取其类型是否与Float64相同
	fmt.Println("kindis float64: ", v.Kind() == reflect.Float64)

	// 获取其值
	fmt.Println("value: ", v.Float())
}
