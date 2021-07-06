package main

import (
	"fmt"
	"math"
)


type FormaGeometrica interface {
	Area() float64
}


func MostrarArea(f FormaGeometrica) {
	area := f.Area()
	fmt.Printf("A área da forma geometrica é %.2f \n", area)
}


type Retangulo struct {
	base float64
	altura float64
}

type Quadrado struct {
	lado float64
}


type Circulo struct {
	raio float64
}


func (r Retangulo) Area() float64{
	return r.base * r.altura
}

func (c Circulo) Area() float64 {
	return math.Pi * math.Pow(c.raio, 2)
}


func main() {
	r := Retangulo{10,15}
	MostrarArea(r)

	c := Circulo{10}
	MostrarArea(c)
}