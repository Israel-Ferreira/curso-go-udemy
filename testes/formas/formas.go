package formas

import (
	"fmt"
	"math"
)

type FormaGeometrica interface {
	area() float64
}

// Quadrado is a Struct to
type Quadrado struct {
	Lado float64
}

func (q Quadrado) Area() float64 {
	return math.Pow(q.Lado, 2)
}

// Retângulo

type Retangulo struct {
	Base   float64
	Altura float64
}

func (r Retangulo) Area() float64 {
	return r.Base * r.Altura
}

// Triângulo
type Triangulo struct {
	Retangulo
}

func (t Triangulo) Area() float64 {
	return (t.Base * t.Altura) / 2
}

func EscreverArea(forma FormaGeometrica)  {
	fmt.Printf("A área da forma é %.2f\n", forma.area())
}


type Circulo struct {
	Raio float64
}


func (c Circulo) Area() float64 {
	return math.Pi * math.Pow(c.Raio,2)
}