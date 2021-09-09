package formas

import (
	"math"
	"testing"
)

func TestFormas(t *testing.T) {
	t.Run("Teste Quadrado", func(t *testing.T) {
		expected := 9.0

		r := Quadrado{3}

		result := r.Area()

		if result != result {
			t.Errorf("Resultado: %.2f, Esperado: %.2f", result, expected)
		}
	})


	t.Run("Teste Retangulo", func(t *testing.T) {
		expected := 6.0

		ret := Retangulo{3, 2}

		result := ret.Area()

		if result != expected {
			t.Errorf("Resultado: %.2f, Esperado: %.2f", result, expected)
		}
	})

	

	t.Run("Teste Triangulo", func(t *testing.T) {

	})

	t.Run("Teste Circulo", func(t *testing.T) {
		c :=  Circulo{2}

		result := c.Area()

		expected := 4 * math.Pi

		if result != expected {
			t.Fatalf("A área recebida %f é diferente da esperada %f", result, expected)
		}
	})
}
