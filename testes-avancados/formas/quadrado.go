package formas

import "math"

type Quadrado struct {
	Lado float64;
}

func (q Quadrado) area() float64 {
	return math.Pow(q.Lado, 2)
}