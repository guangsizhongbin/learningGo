package concise

import (
	"fmt"
	"testing"
)

func Test_Map(t *testing.T) {
	m1 := make(map[string]int)

	m2 := map[string]string{
		"Sam":   "Male",
		"Alice": "Female",
	}

	m1["Tom"] = 18

	fmt.Printf("m1 Map　的值是 %v \n", m1)
	fmt.Printf("m2 Map　的值是 %v \n", m2)
}
