package enderecos

import "strings"

func TipoDeEndereco(enderecos string) string {
	tiposValidos := []string{"rua", "avenida", "estrada", "rodovia"}
	enderecosEmLetraMinuscula := strings.ToLower(enderecos)

	primeiraPalavraDoEndereco := strings.Split(enderecosEmLetraMinuscula, " ")[0]

	enderecoTemUmTipoValido := false
	for _, tipo := range tiposValidos {
		if tipo == primeiraPalavraDoEndereco {
			enderecoTemUmTipoValido = true
		}
	}

	if enderecoTemUmTipoValido {
		return primeiraPalavraDoEndereco
	}

	return "Tipo invalido"
}
