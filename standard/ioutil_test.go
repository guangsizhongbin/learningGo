package standard

import (
	"fmt"
	"io/ioutil"
	"os"
	"testing"
)

// ioutil.ReadAll 会一次性全部读完
func TestReadAll(t *testing.T) {
	f, err := os.Open("reader.txt")
	if err != nil {
		fmt.Println(err)
	}
	defer f.Close()

	b, err := ioutil.ReadAll(f)
	if err != nil {
		fmt.Println("")
	}

	fmt.Println(string(b))
}

// ioutil.ReadDir 读取当前目录下所有的文件
func TestListAll(t *testing.T) {
	fileInfos, err := ioutil.ReadDir(".")
	if err != nil {
		fmt.Println(err)
	}

	for _, info := range fileInfos {
		if info.IsDir() {
			fmt.Println(string(info.Name))
		} else {

		}
	}

}
