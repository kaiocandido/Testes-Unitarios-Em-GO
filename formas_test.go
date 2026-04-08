package main

import (
	"math"
	"testing"
)

func TestAreaRetangulo(t *testing.T) {
	t.Run("Retangulo", func(t *testing.T) {
		ret := Retangulo{Largura: 10, Altura: 12}
		areaEsperada := float64(120)
		AreaRecebida := ret.Area()

		if AreaRecebida != areaEsperada {
			t.Fatalf("Area esperada: %f, Area recebida: %f", areaEsperada, AreaRecebida)
		}

	})
	t.Run("Circulo", func(t *testing.T) {
		cir := Circulo{Raio: 10}
		areaEsperada := float64(math.Pi * 100)
		AreaRecebida := cir.Area()

		if areaEsperada != AreaRecebida {
			t.Errorf("Area esperado: %f, Area recebida %f", areaEsperada, AreaRecebida)
		}
	})
}
