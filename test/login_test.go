package test

import (
	"jwt-todo/estructura"
	"testing"
)

func TestLoginValido(t *testing.T) {
	Username := "username"
	Password := "password"

	if estructura.User1.Username != Username {
		t.Fatal("se esperaba que el usuario fuera invalido")
	}

	if estructura.User1.Password != Password {
		t.Fatal("se esperaba que la contraseña fuera invalida")
	}

}

func TestLoginInvalido(t *testing.T) {
	Username := "user"
	Password := "pass"

	if estructura.User1.Username != Username {
		t.Fatalf("el username: %s es invalido, se esperaba %s", Username, estructura.User1.Username)
	}

	if estructura.User1.Password != Password {
		t.Fatalf("el password: %s es invalido, se esperaba %s", Password, estructura.User1.Password)
	}

}
