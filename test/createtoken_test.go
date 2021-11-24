package test

import (
	"jwt-todo/estructura"
	"testing"
)

func TestCreateTokenInvalido(t *testing.T) {
	idinvalido := uint64(2)
	id := estructura.User1.ID
	if id != idinvalido {
		t.Fatalf("error con el id %v, se esperaba que fuera %v", idinvalido, id)
	}
}

func TestCreateTokenValido(t *testing.T) {
	invalido := uint64(1)
	id := estructura.User1.ID
	if id != invalido {
		t.Fatalf("error con el id %v, se esperaba que fuera %v", invalido, id)
	}
}
