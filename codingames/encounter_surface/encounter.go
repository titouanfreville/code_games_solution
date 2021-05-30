package main

import (
	"fmt"
	"math"
	"os"
	"sort"
)

type (
	Point struct {
		X float64
		Y float64
	}

	Polygon struct {
		Vertexes         []Point
		Segments         []Segment
		ReversedVertexes []Point
		Center           Point
	}

	Segment struct {
		A Point
		B Point

		XFactor float64
		YFactor float64
		EqPoint float64
	}
)

func NewPoint(x, y float64) Point {
	return Point{X: x, Y: y}
}

func (pt Point) OnSegment(s Segment) bool {
	return s.XFactor*pt.X+s.YFactor*pt.Y == s.EqPoint &&
		math.Min(s.A.X, s.B.X) <= pt.X && math.Max(s.A.X, s.B.X) >= pt.X &&
		math.Min(s.A.Y, s.B.Y) <= pt.Y && math.Max(s.A.Y, s.B.Y) >= pt.Y

}

func (pt Point) InPolygon(pol Polygon) (res bool) {
	res = false

	for _, segment := range pol.Segments {
		if (segment.A.Y > pt.Y) != (segment.B.Y > pt.Y) &&
			(pt.X < (segment.A.X-segment.B.X)*(pt.Y-segment.B.Y)/(segment.A.Y-segment.B.Y)+segment.B.X) {
			res = !res
		}
	}

	return res
}

func (pt Point) InList(l []Point) bool {
	for _, p := range l {
		if p == pt {
			return true
		}
	}
	return false
}

func NewPolygon() *Polygon {
	return &Polygon{}
}

func (pol *Polygon) AddVertex(p Point) *Polygon {
	if p.X == 0 {
		p.X = 0
	}

	if !p.InList(pol.Vertexes) {
		pol.Vertexes = append(pol.Vertexes, p)
	}

	return pol
}

func (pol *Polygon) computeCenter() *Polygon {
	var (
		x, y float64
		i    int
		v    Point
	)
	for i, v = range pol.Vertexes {
		x += v.X
		y += v.Y
		i++
	}

	pol.Center = Point{X: x / float64(i), Y: y / float64(i)}
	return pol
}

func (pol *Polygon) sortClockWise() *Polygon {
	sort.Slice(pol.Vertexes, func(i, j int) bool {
		v1 := pol.Vertexes[i]
		v2 := pol.Vertexes[j]
		tan1 := math.Atan2(v1.Y-pol.Center.Y, v1.X-pol.Center.X)
		tan2 := math.Atan2(v2.Y-pol.Center.Y, v2.X-pol.Center.X)
		return tan1 > tan2
	})

	return pol
}

func (pol *Polygon) Build() {
	var p0, pPrev Point

	pol.computeCenter()
	pol.sortClockWise()

	for i, vertex := range pol.Vertexes {
		if i == 0 {
			p0 = vertex
			pPrev = p0
			continue
		}

		pol.Segments = append(pol.Segments, NewSegment(pPrev, vertex))
		pPrev = vertex
	}

	pol.Segments = append(pol.Segments, NewSegment(pPrev, p0))

	for i := len(pol.Vertexes) - 1; i >= 0; i-- {
		pol.ReversedVertexes = append(pol.ReversedVertexes, pol.Vertexes[i])
	}

}

func (pol Polygon) Intersection(pol2 Polygon) Polygon {
	var intersectPoints []Point

	// Compute inner point
	for _, p := range pol.Vertexes {
		if p.InPolygon(pol2) {
			intersectPoints = append(intersectPoints, p)
		}
	}

	for _, p := range pol2.Vertexes {
		if p.InPolygon(pol) {
			intersectPoints = append(intersectPoints, p)
		}
	}

	// Add intersection
	for _, s1 := range pol.Segments {
		for _, s2 := range pol2.Segments {
			if intersectPoint := s1.Intersect(s2); intersectPoint != nil {
				intersectPoints = append(intersectPoints, *intersectPoint)
			}
		}
	}

	debug(intersectPoints)

	res := &Polygon{}

	// Build polygon
	for _, p := range intersectPoints {
		res.AddVertex(p)
	}

	res.Build()

	return *res
}

// Area compute ceiled polygon area
func (pol Polygon) Area() int {
	var x0, y0, xPrev, yPrev, sumX, sumY float64
	sumX = 0
	sumY = 0

	for i, vertex := range pol.ReversedVertexes {
		if i == 0 {
			x0 = vertex.X
			y0 = vertex.Y
			xPrev = x0
			yPrev = y0
			continue
		}

		sumX += xPrev * vertex.Y
		sumY += yPrev * vertex.X

		xPrev = vertex.X
		yPrev = vertex.Y

	}

	sumX += xPrev * y0
	sumY += yPrev * x0

	if sumX > sumY {
		return int(math.Ceil(0.5 * (sumX - sumY)))
	}

	return int(math.Ceil(0.5 * (sumY - sumX)))
}

func NewSegment(a, b Point) Segment {
	xFactor := b.Y - a.Y
	yFactor := a.X - b.X
	eqPoint := xFactor*a.X + yFactor*a.Y
	return Segment{A: a, B: b, XFactor: xFactor, YFactor: yFactor, EqPoint: eqPoint}
}

func (s1 Segment) Intersect(s2 Segment) *Point {
	det := s1.XFactor*s2.YFactor - s2.XFactor*s1.YFactor

	if det == 0 {
		return nil // parallel
	}

	x := (s2.YFactor*s1.EqPoint - s1.YFactor*s2.EqPoint) / det
	y := (s1.XFactor*s2.EqPoint - s2.XFactor*s1.EqPoint) / det
	pt := Point{X: x, Y: y}

	if pt.OnSegment(s1) && pt.OnSegment(s2) {
		return &pt
	}

	return nil
}

func main() {
	var (
		nbVertexArmy1, nbVertexArmy2 int

		x, y float64

		army1 = NewPolygon()
		army2 = NewPolygon()
	)

	fmt.Scan(&nbVertexArmy1)
	fmt.Scan(&nbVertexArmy2)

	for i := 0; i < nbVertexArmy1; i++ {
		fmt.Scan(&x, &y)
		army1.AddVertex(NewPoint(x, y))
	}

	for i := 0; i < nbVertexArmy2; i++ {
		fmt.Scan(&x, &y)
		army2.AddVertex(NewPoint(x, y))
	}

	army1.Build()
	army2.Build()

	intersect := army1.Intersection(*army2)
	debugObjects(intersect)
	fmt.Println(intersect.Area()) // Write answer to stdout
}

func debug(values ...interface{}) {
	fmt.Fprintln(os.Stderr, values...)
}

func debugObjects(values ...interface{}) {
	for _, v := range values {
		fmt.Fprintf(os.Stderr, "%#v\n", v)
	}
}
