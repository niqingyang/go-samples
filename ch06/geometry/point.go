package geometry

import (
	"math"
)

type Point struct {
	X, Y float64
}

func Distance(p, q Point) float64 {
	// 计算直角三角形的斜边长
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p Point) Distance(q Point) float64 {
	// 计算直角三角形的斜边长
	return math.Hypot(q.X-p.X, q.Y-p.Y)
}

func (p *Point) ScaleBy(factor float64)  {
	p.X *= factor
	p.Y *= factor
}

// Path 是连接多个点的直线段
type Path []Point

func (path Path) Distance() float64 {
	sum := 0.0
	for i := range path {
		if i > 0 {
			sum += path[i-1].Distance(path[i])
		}
	}
	return sum
}
