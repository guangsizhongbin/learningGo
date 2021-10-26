package concise

import (
	"fmt"
	"strconv"
	"testing"
)

type Building struct {
	name  string
	age   int
	phone string
}

func (b Building) String() string {
	return "?" + b.name + "-" + strconv.Itoa(b.age) + "years" + b.phone + "?"
}

// 如果需要某个类型能被fmt包以特殊的格式输出，你就必须实现Stringer这个接口
// 如果没有实现这个接口，fmt将以默认的方式输出
func TestCustomString(t *testing.T) {
	Bob := Building{"Bob", 39, "1234"}

	//This Human is : {Bob 39 1234}
	//This Human is : ?Bob-39years1234?
	fmt.Println("This Human is :", Bob)

}
