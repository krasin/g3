package g3

import "math"

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

var ZeroNode = Node{0, 0, 0}

// AdjNodes26 is the list of all 26 adjacent nodes of ZeroNode.
var AdjNodes26 = adjNodes26()

var AdjNodes6 = []Node{
	{0, 0, -1},
	{0, 0, 1},
	{0, -1, 0},
	{0, 1, 0},
	{-1, 0, 0},
	{1, 0, 0},
}

func adjNodes26() (res []Node) {
	for i := 0; i < 27; i++ {
		cur := Node{(i % 3) - 1, ((i / 3) % 3) - 1, ((i / 9) % 3) - 1}
		if !cur.IsZero() {
			res = append(res, cur)
		}
	}
	return
}

// Add returns a sum of node and another node.
func (node Node) Add(another Node) Node {
	return Node{
		node[0] + another[0],
		node[1] + another[1],
		node[2] + another[2],
	}
}

// Sub returns a difference between node and another node.
func (node Node) Sub(another Node) Node {
	return Node{
		node[0] - another[0],
		node[1] - another[1],
		node[2] - another[2],
	}
}

// Point converts Node to Point.
func (node Node) Point() Point {
	return Point{
		float64(node[0]),
		float64(node[1]),
		float64(node[2]),
	}
}

// Vector converts Node to Vector.
func (node Node) Vector() Vector {
	return Vector{
		float64(node[0]),
		float64(node[1]),
		float64(node[2]),
	}
}

// IsZero returns if node is { 0, 0, 0 }
func (node Node) IsZero() bool {
	return node == ZeroNode
}

// L2 calculates L2 norm of vector (aka length)
func (vec Vector) L2() float64 {
	return math.Sqrt(vec[0]*vec[0] + vec[1]*vec[1] + vec[2]*vec[2])
}

// Normalize returns a vector that points to the same direction but has L2 norm equals to 1.
// It will panic on Zero vector and will be out of precision for small vectors.
func (vec Vector) Normalize() Vector {
	l2 := vec.L2()
	return Vector{
		vec[0] / l2,
		vec[1] / l2,
		vec[2] / l2,
	}
}

func (g Grid) At(node Node) Point {
	return Point{
		g.P0[0] + g.H*float64(node[0]),
		g.P0[1] + g.H*float64(node[1]),
		g.P0[2] + g.H*float64(node[2]),
	}
}

func (g Grid) Node(p Point) Node {
	delta := p.Sub(g.P0)
	return Node{
		int(delta[0] / g.H),
		int(delta[1] / g.H),
		int(delta[2] / g.H),
	}
}

func (p Point) Sub(p2 Point) Point {
	return Point{
		p[0] - p2[0],
		p[1] - p2[1],
		p[2] - p2[2],
	}
}
