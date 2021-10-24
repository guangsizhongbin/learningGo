package concise

import (
	"fmt"
	"testing"
)

func Test_Array(t *testing.T) {
	arr := [5]int{1, 2, 3, 4, 5}

	for i := 0; i < len(arr); i++ {
		fmt.Printf("i的值为 %d \n", arr[i])
	}
}

func Test_slice(t *testing.T) {
	slice1 := make([]float32, 0)    // 长度为0的切片
	slice2 := make([]float32, 3, 5) // 长度为0的切片
	fmt.Println(len(slice2), cap(slice2))
	fmt.Println(len(slice1), cap(slice1))



}
