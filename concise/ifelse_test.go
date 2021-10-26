package concise

import (
	"fmt"
	"testing"
)

func TestFor(t *testing.T){
    age := 18
    if age < 18 {
        fmt.Printf("kid")
    } else {
        fmt.Printf("Adult")
    }
}

func TestSimFor(t *testing.T){
    if age := 18; age < 18{
        fmt.Printf("kid")
    } else {
        fmt.Printf("Adult")
    }
}

func TestIfElse(t *testing.T){
    integer := 3

    if integer == 3{
        fmt.Println("The integer is equal to 3")
    } else if integer < 3 {
        fmt.Println("The integer is less than 3")
    } else {
        fmt.Println("The integer is greater than 3")
    }
}