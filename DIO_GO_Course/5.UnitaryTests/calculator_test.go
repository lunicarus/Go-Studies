package main

import "testing"

func TestShouldMultiplyAndPass(t *testing.T) {
	teste := multiply(2, 3)

	if teste != 6 {
		t.Error("Valor Esperado: ", 6, "Valor Retornado: ", teste)
	}
}

func TestShouldResultInZero(t *testing.T) {
	teste := multiply(0, 5)

	if teste != 0 {
		t.Error("Valor Esperado: ", 0, "Valor Retornado: ", teste)
	}
}

func TestShouldAddAndPass(t *testing.T) {
	teste := add(2, 3)

	if teste != 5 {
		t.Error("Valor Esperado: ", 5, "Valor Retornado: ", teste)
	}
}

func TestShouldSubAndPass(t *testing.T) {
	teste := sub(2, 3)

	if teste != -1 {
		t.Error("Valor Esperado: ", -1, "Valor Retornado: ", teste)
	}
}

func TestShouldDivAndPass(t *testing.T) {
	teste, _ := divide(4.5, 3)

	if teste != 1.5 {
		t.Error("Valor Esperado: ", 1.5, "Valor Retornado: ", teste)
	}
}
func TestShouldDiv0AndFail(t *testing.T) {
	teste, err := divide(0, 3)

	if teste != 0 && err != nil {
		t.Error("Falha ao Gerar Erro de Divisão por 0")
	}
}
func TestShouldFailWhenGiving0Values(t *testing.T) {
	teste, err := divide()

	if teste != 0 && err != nil {
		t.Error("Falha ao Gerar Erro de Divisão por 0")
	}
}
