package main

import (
	"fmt"
	"introduction-tests/enderecos"
	"testing"
)

func AssertHelper(result, expected interface{}, t *testing.T) {
	t.Helper()

	if result != expected {
		t.Errorf("Result: %s, Expected: %s", result, expected)
	}

}

func TestTipoEndereco(t *testing.T) {

	t.Run("Avenida", func(t *testing.T) {
		endereco := "Avenida Paulista, 1890"
		result := enderecos.TipoDeEndereco(endereco)
		expected := "Avenida"

		AssertHelper(result, expected, t)
	})

	t.Run("Rua", func(t *testing.T) {
		endereco := "Rua Augusta, 2356"

		result := enderecos.TipoDeEndereco(endereco)
		expected := "Rua"

		AssertHelper(result, expected, t)
	})

}

func TestContains(t *testing.T) {
	quote := "Quick brown fox jumps over the lazy dog"
	word := "fox"

	result := enderecos.Contains(quote, word)
	expected := true

	AssertHelper(result, expected, t)
}

func ExampleContains() {
	endereco := "Avenida Cruzeiro do Sul, 3456"
	result := enderecos.TipoDeEndereco(endereco)

	fmt.Println(result)

	// Output: Avenida
}
