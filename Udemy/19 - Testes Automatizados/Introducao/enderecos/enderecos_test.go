package enderecos_test

import (
	. "introducao-testes/enderecos"
	"testing"
)

type cenarioDeTeste struct {
	enderecoInserido string
	retornoEsperado  string
}

func TestTipeDeEndereco(t *testing.T) {

	cenariosDeTestes := []cenarioDeTeste{
		{"rua ABC", "rua"},
		{"avenida ABC", "avenida"},
		{"rodovia ABC", "rodovia"},
		{"Praça ABC", "Tipo invalido"},
		{"estrada ABC", "estrada"},
		{"", "Tipo invalido"},
	}

	for _, cenario := range cenariosDeTestes {
		TipoDeEnderecoRecebido := TipoDeEndereco(cenario.enderecoInserido)
		if TipoDeEnderecoRecebido != cenario.retornoEsperado {
			t.Errorf("O tipo recebido %s é diferente do esperado %s",
				TipoDeEnderecoRecebido,
				cenario.retornoEsperado)
		}
	}

}
