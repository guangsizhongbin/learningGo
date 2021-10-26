package concise

import (
	"fmt"
	"testing"
)

func TestForsum(t *testing.T) {
	sum := 0
	// for i := 0; i < 50; i++ {
	// 	sum += i
	// }
	for i := 0; i < 100; i++ {
		if sum > 50 {
			break
		}
	sum += i
	}
	fmt.Printf("from 0 sum to 50, %d", sum)
}

// for 相当于　while
func TestForEqualWhile(t *testing.T){
	// 1.　忽略1,3
	sum := 1
	for ; sum < 1000;{
		sum += sum
	}
	fmt.Println(sum)

	// 2. 忽略1,3 并且 ; 也省略
	newSum := 1
	for newSum < 1000 {
		newSum += newSum
	}
	fmt.Println(newSum)
}

// for ＋　break
func TestForWithBreak(t *testing.T){
	for index := 10; index > 0; index--{
		if index == 5{
			break
		}
		fmt.Println(index)
	}
}

// for + continue
func TestForWithContinue(t *testing.T){
	for index := 10; index > 0; index--{
		if index == 5{
			continue
		}
		fmt.Println(index)
	}
}
