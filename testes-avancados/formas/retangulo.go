package formas

type Retangulo struct {
	Base float64;
	Altura float64;
}


func (r Retangulo) area() float64 {
	return r.Base * r.Altura
}
