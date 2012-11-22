package g3

// Point is a point in 3d space.
type Point [3]float64

// Vector is a vector in 3d space.
type Vector [3]float64

// Triangle is a triangle in 3d space with a normal defined by the order of vertices.
// See http://www.opengl.org/wiki/Calculating_a_Surface_Normal
type Triangle [3]Point

// Node is a point in 3d space with integer coordinates. It also might be a node in a Grid.
type Node [3]int

// Grid is a uniform grid NxNxN with step H in the cube with smallest point P0.
type Grid struct {
	P0 Point
	H  float64
	N  int
}