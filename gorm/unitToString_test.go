package gorm

import (
	"fmt"
	"strconv"
	"testing"
)

func TestUintToString(t *testing.T) {
	var n uint = 123
	var s string = strconv.FormatUint(uint64(n), 10)
	fmt.Printf("s = %s \n", s)
}
