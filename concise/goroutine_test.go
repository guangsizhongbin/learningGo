package concise

import (
	"fmt"
	"sync"
	"testing"
	"time"
)

func TestGoroutine(t *testing.T) {
	go count("sleep123")
	count("apple")

}

func TestGoroutineWithTwoGo(t *testing.T) {
	go count("sleep123")
	go count("apple")
	// 使用go后, 这个协程会在后面执行，直接进入下一条命令, 但两个协程都执行完就没有了
}

func TestGoroutineWithTwoGoAndSleep(t *testing.T) {
	go count("sleep123")
	go count("apple")
	time.Sleep(time.Millisecond * 500)
}

func TestGoroutineWithTwoGoAndScanln(t *testing.T) {
	go count("sleep123")
	go count("apple")

	// 等待用户输入
	fmt.Scanln()
}

func TestGoroutineWithSync(t *testing.T){
	var wg sync.WaitGroup
	wg.Add(1)

	go func(){
		count("sheep")
		wg.Done()
	}()
	wg.Wait()
}

func count(something string) {
	for i := 0; i < 100; i++ {
		fmt.Println(something)
	}
	time.Sleep(time.Microsecond * 500)
}
