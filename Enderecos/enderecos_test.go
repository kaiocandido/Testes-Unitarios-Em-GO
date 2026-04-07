// Teste de Unidade para a função TipoDeEndereco do pacote enderecos
package enderecos

import "testing"

func TestTipoDeEndereco(t *testing.T) {
	enderecoParaTeste := "Rua das Flores"
	TipoDeEnderecoEsperado := "Rua"
	tipoDeEnderecoRecebido := TipoDeEndereco(enderecoParaTeste)

	if tipoDeEnderecoRecebido != TipoDeEnderecoEsperado {
		t.Error("O tipo recebido é diferente do esperado!")
	}
}
