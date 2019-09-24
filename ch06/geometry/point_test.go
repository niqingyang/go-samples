package geometry

import (
	"fmt"
	"testing"
)

// 计算两点间的距离
func TestPoint(t *testing.T) {
	p := Point{1, 2}
	q := Point{4, 6}

	fmt.Println(Distance(p, q)) // 5，函数调用
	fmt.Println(p.Distance(q))  // 5，方法调用
}

func TestPath(t *testing.T) {
	perim := Path{
		{1, 1},
		{5, 1},
		{5, 4},
		{1, 1},
	}

	fmt.Println(perim.Distance()) // 12
}

func TestScaleBy(t *testing.T) {
	r := &Point{1, 2}
	r.ScaleBy(2)
	fmt.Println(*r) // {2 4}

	// 或者

	{
		p := Point{1,2}
		p.ScaleBy(2)
		fmt.Println(p) // {2 4}
	}

	// 或者

	p := Point{1,2}
	pptr := &p
	pptr.ScaleBy(2)
	fmt.Println(p) // {2 4}

	// 或者

	{
		p := Point{1,2}
		(&p).ScaleBy(2)
		fmt.Println(p) // {2 4}
	}
}
