package nats

import (
	"fmt"
	"testing"

	"github.com/nats-io/nats.go"
)

func TestNatsServer(t *testing.T) {
	nc, _ := nats.Connect(nats.DefaultURL)

	// 简单的异步 subscriber
	// nc.Subscribe("foo", func(m *nats.Msg) {
	// 	fmt.Printf("Received a message: %s\n", string(m.Data))
	// })

	// 回复 Subscriber
	nc.Subscribe("foo", func(m *nats.Msg) {
		m.Respond([]byte("answer is 42"))
	})

	// 设置同步的 Subscriber

	// 键盘输入, 阻塞线程
	fmt.Scanln()
}
