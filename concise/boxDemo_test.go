package concise

import (
	"fmt"
	"testing"
)

const (
	WHITE = iota
	BLACK
	BLUE
	RED
	YELLOW
)

type Color byte

type Box struct {
	width, height, depth float64
	color                Color
}

type BoxList []Box // a slice of boxes

func (b Box) Volume() float64 {
	return b.width * b.height * b.depth
}


func (b *Box) SetColor(c Color) {
	// *.color = c 也是对的，　Go语言知道你要通过指针去获取这个值
    b.color = c
}

func (b1 BoxList) BiggestsColor() Color {
	v := 0.00
	k := Color(WHITE)
	for _, b := range b1 {
		if b.Volume() > v {
			v = b.Volume()
			k = b.color
		}
	}
	return k
}

func (b1 BoxList) PaintItBlack() {
	for i, _ := range b1 {
        // &b1[i].SetColor(BLACK)也是对的
        // Go语言知道receiver是指针，它自动转换了
		b1[i].SetColor(BLACK)
	}
}

func (c Color) String() string {
	strings := []string{"WHITE", "BLACK", "BLUE", "RED", "YELLOW"}
	return strings[c]
}

func TestBoxDemo(t *testing.T) {
	boxes := BoxList{
		Box{4, 4, 4, RED},
		Box{10, 10, 1, YELLOW},
		Box{1, 1, 20, BLACK},
		Box{10, 10, 20, BLUE},
		Box{10, 30, 1, WHITE},
		Box{20, 20, 20, YELLOW},
	}

    fmt.Printf("We have %d boxes in our set \n", len(boxes))
    fmt.Println("The volume of the first one is ", boxes[0].Volume(), "cm2")
    fmt.Println("The color of the last one is", boxes[len(boxes) - 1])
}
