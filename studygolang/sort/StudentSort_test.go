package sort

import (
	"fmt"
	"sort"
	"testing"
)

type StuScore struct {
	name  string // 姓名
	score int    // 成绩
}

type StuScores []StuScore

// 长度
func (s StuScores) Len() int {
	return len(s)
}

// 成绩将有低到高排序
func (s StuScores) Less(i, j int) bool {
	return s[i].score > s[j].score
}

// Swap()
func (s StuScores) Swap(i, j int) {
	s[i], s[j] = s[j], s[i]
}

func TestMain(t *testing.T) {
	stus := StuScores{
		{"alan", 95},
		{"hikerell", 91},
		{"acmfly", 96},
		{"leao", 90},
	}

	// 打印未排序
	fmt.Println("Default:\n\t", stus)

	// StuScores 已经实现了 sort.Interface 接口, 所以可以调用 Sort 函数进行排序
	sort.Sort(stus)

	// 判断是否已经排好顺序,将会打印  true
	fmt.Println("Is Sorted?\n\t", sort.IsSorted(stus))

	// 打印排序后的 stus 数据
	fmt.Println("Sorted:\n\t", stus)

	fmt.Println(sort.Reverse(sort.IntSlice(stus))

}
