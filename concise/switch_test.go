package concise

import (
	"fmt"
	"testing"
)

type Gender int8

const (
	MALE   Gender = 1
	FEMALE Gender = 2
)

func Test_switch(t *testing.T) {

	gender := MALE
	switchWithBreak(gender)

	switchfallthrough(gender)

}

func switchWithBreak(gender Gender) {
	switch gender {
	case FEMALE:
		fmt.Println("female")
	case MALE:
		fmt.Println("male")
	default:
		fmt.Println("unknown")
	}
}

func switchfallthrough(gender Gender) {
	switch gender {
	case FEMALE:
		fmt.Println("female")
		fallthrough
	case MALE:
		fmt.Println("female")
		fallthrough
	default:
		fmt.Println("unknown")
	}
}
