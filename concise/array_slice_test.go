package concise

import (
	"fmt"
	"testing"
)

// sliece 和数组的声明数组时，方括号内写明了数组的长度或使用...自动计算长度, 而声明slice时, 方括号内没有任何字符
func TestArray(t *testing.T) {
	arr := [5]int{1, 2, 3, 4, 5}

	for i := 0; i < len(arr); i++ {
		fmt.Printf("i的值为 %d \n", arr[i])
	}
}

func TestSlice(t *testing.T) {
	slice1 := make([]float32, 0)    // 长度为0的切片
	slice2 := make([]float32, 3, 5) // 长度为0的切片
	fmt.Println(len(slice2), cap(slice2))
	fmt.Println(len(slice1), cap(slice1))

	slice := []byte{'a', 'b', 'c', 'd'}
	for _, v := range slice {
		fmt.Printf("其保存的数值是%d, 对应的字符是%c \n", v, v)
	}

}

func TestSimpleCreateArray(t *testing.T) {
	b := [10]int{1, 2, 3} //声明了一个长度为10的int数组，其中前三个元素初始化为1, 2, 3, 默认为0

	c := [...]int{4, 5, 6} //可以省略长度而采用 `...` 的方式，　Go语言会自动根据元素个数来计算长度

	// for i := 0; i < len(b); i++ {
	// 	fmt.Printf("%d ", b[i])
	// }
	// fmt.Println()

	// for i := 0; i < len(c); i++ {
	// 	fmt.Printf("%d ", c[i])
	// }

	for _, v := range b {
		fmt.Printf("%d ", v)
	}
	fmt.Println(" ")
	for _, v := range c {
		fmt.Printf("%d ", v)
	}
}

// 定义二维数组，并遍历二维数组
func TestTwoDimesionalArray(t *testing.T) {
	doubleArray := [2][4]int{[4]int{1, 2, 3, 4}, [4]int{5, 6, 7, 8}}
	easyArray := [2][4]int{{1, 2, 3, 4}, {5, 6, 7, 8}}

	for i, v := range doubleArray {
		fmt.Print(i, ":")
		for _, value := range v {
			fmt.Print(value)
		}
		fmt.Println()
	}

	for i, v := range easyArray {
		fmt.Print(i, ":")
		for _, value := range v {
			fmt.Print(value)
		}
		fmt.Println()
	}
}

func TestSliceLenAndCap(t *testing.T){
	var array = [10]byte{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j'}

	aSlice := array[:3]
	fmt.Printf("aSlice当前的值%v, 其长度为%d, 容量为%d\n", aSlice, len(aSlice), cap(aSlice))

	bSlice := append(aSlice, 1)
	fmt.Printf("bSlice当前的值%v, 其长度为%d, 容量为%d\n", bSlice, len(bSlice), cap(bSlice))

	cSlice := array[3:]
	fmt.Printf("cSlice当前的值%v, 其长度为%d, 容量为%d\n", cSlice, len(cSlice), cap(cSlice))

	// 当slice容量不够时，会触发slice的自动扩容
	dSlice := append(cSlice, 1)
	fmt.Printf("dSlice当前的值%v, 其长度为%d, 容量为%d\n", dSlice, len(dSlice), cap(dSlice))
}
