package enderecos

import "strings"

// TipoDeEndereco recebe um endereço como string e retorna o tipo de endereço (Rua, Avenida, Alameda, Travessa) ou uma mensagem de erro se o tipo for inválido.
func TipoDeEndereco(endereco string) string {
	tiposValidos := []string{"Rua", "Avenida", "Alameda", "Travessa"}
	enderecoFormatado := strings.ToLower(endereco)
	primeiraPalavra := strings.Split(enderecoFormatado, " ")[0]
	enderecoValido := false

	for _, tipo := range tiposValidos {
		if strings.ToLower(tipo) == primeiraPalavra {
			enderecoValido = true
			break
		}
	}

	if enderecoValido {
		enderecoFormatadoSaida := strings.ToTitle(primeiraPalavra)
		return enderecoFormatadoSaida
	}

	return "Tipo de endereço inválido"
}
