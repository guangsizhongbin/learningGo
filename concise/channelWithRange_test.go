package concise

import (
	"fmt"
	"testing"
)

// 应该在生产者的地方关闭 channel, 而不是消费的地方去关闭它，这样容易引起panic
// channel 不像文件之类需要经常去关闭，只有当你确实没有任何数据发送了，或你想显式的结束range循环之类的操作
func fibonacci(n int, c chan int) {
	x, y := 1, 1
	for i := 0; i < n; i++ {
		c <- x
		x, y = y, x+y
	}
	close(c)
}

func TestChannelWithRange(t *testing.T) {
	c := make(chan int, 10)
	go fibonacci(cap(c), c)
	for i := range c {
		fmt.Println(i)
	}
}
