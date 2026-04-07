// Teste de Unidade para a função TipoDeEndereco do pacote enderecos
package enderecos

import (
	"strings"
	"testing"
)

func TestTipoDeEndereco(t *testing.T) {
	enderecoParaTeste := "Rua das Flores"
	TipoDeEnderecoEsperado := "Rua"
	tipoDeEnderecoRecebido := strings.ToTitle(TipoDeEndereco(enderecoParaTeste))

	if tipoDeEnderecoRecebido != TipoDeEnderecoEsperado {
		t.Errorf("O tipo recebido é diferente do esperado! Recebido: %s, Esperado: %s", tipoDeEnderecoRecebido, TipoDeEnderecoEsperado)
	}
}
