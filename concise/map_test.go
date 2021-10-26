package concise

import (
	"fmt"
	"testing"
)

func TestMap(t *testing.T) {
	m1 := make(map[string]int)

	m2 := map[string]string{
		"Sam":   "Male",
		"Alice": "Female",
	}

	m1["Tom"] = 18

	fmt.Printf("m1 Map　的值是 %v \n", m1)
	fmt.Printf("m2 Map　的值是 %v \n", m2)
}

func TestMapElemetIfExist(t *testing.T){
	// 初始化一个字典
	rating := map[string]float32 {"c": 5, "Go": 4.5, "Python": 4.5, "C++": 2}
	fmt.Printf("当前rating的长度是 %d\n", len(rating))

	// 判断是否有值
	csharpRating, ok := rating["C#"]
	if ok {
		fmt.Printf("C# is in the map and its rating is %f \n", csharpRating)
	} else {
		fmt.Printf("We have no rating associated with C# in the map \n")
	}
}

func TestMapDeleteElement(t *testing.T){
	// 初始化一个字典
	rating := map[string]float32 {"c": 5, "Go": 4.5, "Python": 4.5, "C++": 2}
	fmt.Printf("当前rating的长度是 %d\n", len(rating))

	delete(rating, "c")
	fmt.Printf("当前rating的长度是 %d\n", len(rating))
	fmt.Printf("当前map是 %v\n", rating)
}
