package concise

import (
	"fmt"
	"testing"
)

func Test_for(t *testing.T){
    age := 18
    if age < 18 {
        fmt.Printf("kid")
    } else {
        fmt.Printf("Adult")
    }
}

func Test_sim_for(t *testing.T){
    if age := 18; age < 18{
        fmt.Printf("kid")
    } else {
        fmt.Printf("Adult")
    }
}
