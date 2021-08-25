package main

import (
	"fmt"
	"introduction-tests/enderecos"
	"testing"
)


type Scenario struct {
	givenAddress string
	expectedAddress string
}

func AssertHelper(result, expected interface{}, t *testing.T) {
	t.Helper()

	if result != expected {
		t.Errorf("Result: %s, Expected: %s", result, expected)
	}

}



func TestTipoEnderecoWithScenarios(t *testing.T){
	scenarios := []Scenario{
		{givenAddress: "Rua ABC", expectedAddress: "Rua"},
		{givenAddress: "Avenida Brigadeiro Faria Lima, 2447", expectedAddress:"Avenida"},
		{givenAddress: "Praça da Sé", expectedAddress: "Tipo Inválido"},
		{givenAddress: "Jardim Botânico de Jundiaí", expectedAddress: "Tipo Inválido"},
		{givenAddress: "Rodovia dos Imigrantes, Km. 23", expectedAddress: "Rodovia"},
		{givenAddress: "Estrada do M'boi Mirim, km 12", expectedAddress: "Estrada" },
	}


	for _, scenario := range scenarios {
		result := enderecos.TipoDeEndereco(scenario.givenAddress)
		AssertHelper(result, scenario.expectedAddress, t)
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


	t.Run("RUA CAPITALIZED", func(t *testing.T) {
		endereco := "RUA DOS BOBOS"
		result := enderecos.TipoDeEndereco(endereco)

		expected := "Rua"

		AssertHelper(result, expected, t)
	})


	t.Run("Rodovia Anhaguera", func(t *testing.T) {
		endereco := "Rodovia Anhanguera, Km 60"

		result := enderecos.TipoDeEndereco(endereco)
		expected := "Rodovia"

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
