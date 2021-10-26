package concise

import (
	"fmt"
	"testing"
	"time"
)

// 必须使用make创建channel
func TestChannel(t *testing.T) {
	c := make(chan string)
	go countSomething("sheep", c)

	for {
		msg, open := <-c

		if !open {
			break
		}

		fmt.Println(msg)
	}
}

func TestChannelWithSugar(t *testing.T) {
	c := make(chan string)
	go countSomething("sheep", c)

	for msg := range c {
		fmt.Println(msg)
	}
}

func countSomething(thing string, c chan string) {
	for i := 1; i <= 5; i++ {
		c <- thing
		time.Sleep(time.Millisecond * 500)
	}

	close(c)
}

// Buffer channel
func TestBufferChannel(t *testing.T) {
	c := make(chan string, 2)
	c <- "hello"
	c <- "world"

	msg := <-c
	fmt.Println(msg)

	msg = <-c
	fmt.Println(msg)
}


