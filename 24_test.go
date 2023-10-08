package main

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"math"
	"testing"
)

// Разработать программу нахождения расстояния между двумя точками, которые
// представлены в виде структуры Point с инкапсулированными параметрами x,y и конструктором.

type Point struct {
	x, y int
}

func NewPoint(x, y int) *Point {
	return &Point{x, y}
}
func (p *Point) CalculateDistance(p2 *Point) float64 {
	xdelta := p.x - p2.x
	ydelta := p.y - p2.y
	return math.Sqrt(float64(xdelta*xdelta + ydelta*ydelta))
}

func Test24(t *testing.T) {
	p1 := NewPoint(1, 1)
	p2 := NewPoint(5, 4)

	fmt.Println(p1.CalculateDistance(p2))

	// Ожидается что расстояние = 5, и что расстояние p1-p2 == p2-p1
	assert.Equal(t, float64(5), p1.CalculateDistance(p2))
	assert.Equal(t, p2.CalculateDistance(p1), p1.CalculateDistance(p2))
}
