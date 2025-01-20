package main

import "testing"

func TestSoma(t *testing.T) {
	teste := soma(3, 2, 1)
	resultado := 6
	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}

func TestMultiplica(t *testing.T) {
	teste := multiplica(10, 10)
	resultado := 100
	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}

func TestSubtrai(t *testing.T) {
	teste := subtrai(10, 5)
	resultado := 5
	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}

func TestDivide(t *testing.T) {
	teste := divide(20, 2)
	resultado := 10
	if teste != resultado {
		t.Error("Valor esperado:", resultado, "Valor retornado:", teste)
	}
}

func TestDivideByZero(t *testing.T) {
	teste := divide(20, 0)
	resultado := 0
	if teste != resultado {
		t.Error("Divisão por zero deveria retornar 0. Valor retornado:", teste)
	}
}
