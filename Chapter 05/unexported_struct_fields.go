package main

type ImmutablePoint struct {
    x int
    y int
}

func NewImmutablePoint(x, y int) *ImmutablePoint {
    return &ImmutablePoint{x: x, y: y}
}

func (p *ImmutablePoint) X() int {
    return p.x
}

func (p *ImmutablePoint) Y() int {
    return p.y
}

func main() {
    point := NewImmutablePoint(10, 20)
    fmt.Println(point.X(), point.Y())
    // point.x = 15 // This line would result in a compilation error
}
