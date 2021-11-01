package standard

import (
	"fmt"
	"io"
	"os"
	"testing"
)

// 基本的IO接口
func TestReaderExample(t *testing.T) {
	// 先打开文件
	file, err := os.Open("reader.txt")
	if err != nil {
		fmt.Println("打开文件 reader.txt 错误:", err)
	}
	// defer 放在错误检查之后
	defer file.Close()

	// 读取文件
	data, err := ReadFrom(file, 12)
	if err != nil {
		fmt.Println("读取文件失败", err)
	}
	fmt.Println(string(data))

}

func ReadFrom(reader io.Reader, num int) ([]byte, error) {
	p := make([]byte, num)
	n, err := reader.Read(p)
	if n > 0 {
		return p[:n], nil
	}
	return p[:n], err

}
