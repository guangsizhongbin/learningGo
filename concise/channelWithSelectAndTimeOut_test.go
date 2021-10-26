package concise

import (
	"testing"
	"time"
)

func TestChannelWithSelectAndTimeOut(t *testing.T) {
	c := make(chan int)
	o := make(chan bool)
	go func() {
		for {
			select {
			case v := <-c:
				println(v)
			case <-time.After(5 * time.Second):
				println("timeout")
				o <- true
			}
		}
	}()
	<-o
}
