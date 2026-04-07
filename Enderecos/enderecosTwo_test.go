package enderecos

import "testing"

// Teste de Unidade para a função TipoDeEndereco do pacote enderecos
type cenarioTest struct {
	enderecoParaTeste      string
	TipoDeEnderecoEsperado string
}

// TipoDeEndereco recebe um endereço como string e retorna o tipo de endereço (Rua, Avenida, Alameda, Travessa) ou uma mensagem de erro se o tipo for inválido.
func TipoDeEnderecoTwo(t *testing.T) string {
	cenariosDeTest := []cenarioTest{
		{enderecoParaTeste: "Rua das Flores", TipoDeEnderecoEsperado: "Rua"},
		{enderecoParaTeste: "Avenida Paulista", TipoDeEnderecoEsperado: "Avenida"},
		{enderecoParaTeste: "Alameda Santos", TipoDeEnderecoEsperado: "Alameda"},
		{enderecoParaTeste: "Travessa do Comércio", TipoDeEnderecoEsperado: "Travessa"},
	}

	for _, cenario := range cenariosDeTest {
		tipoDeEnderecoRecebido := TipoDeEndereco(cenario.enderecoParaTeste)
		if tipoDeEnderecoRecebido != cenario.TipoDeEnderecoEsperado {
			t.Errorf("O tipo recebido é diferente do esperado! Recebido: %s, Esperado: %s", tipoDeEnderecoRecebido, cenario.TipoDeEnderecoEsperado)
		}
	}

	return "Tipo de endereço inválido"
}
