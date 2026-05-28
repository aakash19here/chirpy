package auth

import (
	"testing"
)

const password = "super-secret"

func TestHashing(t *testing.T) {
	hash, err := HashPassword(password)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if hash == "" {
		t.Fatal("expected hash to not be empty")
	}

	if hash == password {
		t.Fatal("hash should not equal plain password")
	}
}

func TestCheckPasswordHash_Valid(t *testing.T) {
	hash, err := HashPassword(password)

	if err != nil {
		t.Fatalf("failed to hash password: %v", err)
	}

	match, err := CheckPasswordHash(password, hash)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if !match {
		t.Fatal("expected password to match hash")
	}
}

func TestCheckPasswordHash_InValid(t *testing.T) {
	hash, err := HashPassword(password)

	if err != nil {
		t.Fatalf("failed to hash password: %v", err)
	}

	match, err := CheckPasswordHash("wrong-password", hash)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}
	if match {
		t.Fatal("expected password to not match hash")
	}
}
