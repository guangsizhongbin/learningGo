package nats

import (
	"log"
	"testing"

	"github.com/nats-io/nats.go"
)

func TestNatConn(t *testing.T) {
	// 使用nats的默认连接, 单个服务器
	// nc, err := nats.Connect(nats.DefaultURL)
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer nc.Close()

	// nats 使用集群
	// servers := []string{"nats://127.0.0.1:1222", "nats://127.0.0.1:1223", "nats://127.0.0.1:1224"}
	// nc, err := nats.Connect(strings.Join(servers, ",")) // 不同的服务器之间用,分隔
	// if err != nil {
	// 	log.Fatal(err)
	// }
	// defer nc.Close()

	// 获取连接状态
	nc, err := nats.Connect(nats.DefaultURL, nats.Name("API Example"))
	if err != nil {
		log.Fatal(err)
	}
	defer nc.Close()

	getStatusTxt := func(nc *nats.Conn) string {
		switch nc.Status() {
		case nats.CONNECTED:
			return "Connected"
		case nats.CLOSED:
			return "Closed"
		default:
			return "Other"
		}
	}

	log.Printf("The connection is %v\n", getStatusTxt(nc))

	nc.Close()

	log.Printf("The connection is %v\n", getStatusTxt(nc))

	// 设置nats的名字
	// Simpe Publisher
	// nc.Publish("foo", []byte("Hello World"))
}
