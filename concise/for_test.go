package concise

import (
	"fmt"
	"testing"
)

func Test_forsum(t *testing.T) {
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
