package concise

import (
	"fmt"
	"testing"
)

// 当存在多个channel时，可以通过select 监听channel上的数据流动
func fibonacciWithSelect(c, quit chan int) {
	x, y := 1, 1
	for {
		select {
		case c <- x:
			x, y = y, x+y
		case <- quit:
			fmt.Println("quit")
			return
		}
	}
}

func TestChannelWithSelect(t *testing.T) {
	c := make(chan int)
	quit := make(chan int)

	go func() {
		for i := 0; i < 10; i++ {
			fmt.Println(<-c)
		}
		quit <- 0
	}()
	fibonacciWithSelect(c, quit)
}
