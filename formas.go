package main

type Retangulo struct {
	Largura float64
	Altura  float64
}

type Circulo struct {
	Raio float64
}

// interface forma geometrica que tem o metodo area() que retorna um float64
// interace serve para definir um contrato, ou seja, qualquer tipo que implemente o metodo area() pode ser considerado uma forma geometrica
type FormaGeometrica interface {
	Area() float64
}

func Area(r FormaGeometrica) {
	area := r.Area()
	println(area)
}

func (r Retangulo) Area() float64 {
	return r.Largura * r.Altura
}

func (c Circulo) Area() float64 {
	return 3.14 * c.Raio * c.Raio
}

func main() {
	r := Retangulo{10, 5}
	c := Circulo{5}
	Area(r)
	Area(c)
}
