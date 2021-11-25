package test

import (
	"testing"

	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	Username string `json:"username"`
	jwt.StandardClaims
}

var jwtKey = []byte("ACCESS_SECRET")

func TestTokenValidValid(t *testing.T) {
	tooken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJlbWFpbCI6Impvc2VAZ21haWwuY29tIiwiZXhwIjoxNjM3Nzg2ODY5LCJsYXNuYW1lIjoiam9zZSIsIm5hbWUiOiJqb3NlIiwicGhvbmUiOiIxMjM0NTY3ODkwIiwidXNlcl9pZCI6MX0.tKEd800bC5wouJmk8Y2BVK3QQHlBOqq-HrgFA_cChTA"
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(tooken, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			t.Fatal("token invalido")
		}
		t.Fatal("token invalido ")
	}
	if !tkn.Valid {
		t.Fatal("token invalido")
	}
}
func TestTokenValidInvalid(t *testing.T) {
	//tooken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJlbWFpbCI6Impvc2VAZ21haWwuY2s9tIiwiZXhwIjoxNjM3NzgyOTE5LCJsYXNuYW1lIjoiam9zZSIsIm5hbWUiOiJqb3NlIiwicGhvbmUiOiIxMjM0NTY3ODkwIiwidXNlcl9pZCI6MX0.pjwgVIWmXva5Mpc_aAViCIAm8PQOfcHtOhMj3MvHzi8"
	tooken := "eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJhdXRob3JpemVkIjp0cnVlLCJlbWFpbCI6Impvc2VAZ21haWwuY29tIiwiZXhwIjoxNjM3Nzg3NzM2LCJsYXNuYW1lIjoiam9zZSIsIm5hbWUiOiJqb3NlIiwicGhvbmUiOiIxMjM0NTY3ODkwIiwidXNlcl9pZCI6MX0.XtVL2Vj5c1tuSK9ENf65L-9W7Jo49ji9FDN4M6pC2aM"
	claims := &Claims{}
	tkn, err := jwt.ParseWithClaims(tooken, claims, func(token *jwt.Token) (interface{}, error) {
		return jwtKey, nil
	})
	if err != nil {
		if err == jwt.ErrSignatureInvalid {
			t.Fatal("token invalido")
		}
		t.Fatal("token invalido ")
	}
	if !tkn.Valid {
		t.Fatal("token invalido")
	}
}
