package gorm

import (
	"encoding/json"
	"fmt"
	"testing"
)

type Post struct {
	Id      int
	Content string
	Author  string
}

func ObjectToString(post Post) ([]byte, error) {
	byte, err := json.Marshal(post)
	if err != nil {
		fmt.Println(nil)
		return byte, nil
	}
	return byte, nil
}

func StringToObject(string []byte) (*Post, error) {
	// json()
	var post Post
	err := json.Unmarshal(string, &post)
	if err != nil {
		return nil, err
	}
	return &post, nil
}

func TestJsonObject(t *testing.T) {
	post1 := Post{1, "Hello World", "userA"}
	b, err := ObjectToString(post1)
	if err != nil {
		fmt.Println("序列化失败")
	}

	fmt.Printf("序列化后的值为: %s\n", string(b))

	p, err := StringToObject(b)
	if err != nil {
		fmt.Println("反序列化失败")
	}
	fmt.Printf("post 的id : %d, author : %s, content: %s ", p.Id, p.Author, p.Content)
}
