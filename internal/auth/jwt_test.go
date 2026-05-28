package auth

import (
	"testing"
	"time"

	"github.com/google/uuid"
)

const userId = "123e4567-e89b-12d3-a456-426614174000"
const tokenSecret = "my-super-secret"
const expiresIn = 1 * time.Hour

func TestMakeJWT(t *testing.T) {
	parsedUUID, err := uuid.Parse(userId)
	if err != nil {
		t.Fatalf("failed to parse test UUID: %v", err)
	}

	token, err := MakeJWT(parsedUUID, tokenSecret, expiresIn)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	u, err := ValidateJWT(token, tokenSecret)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if u != parsedUUID {
		t.Fatal("expected user id to match")
	}
}

func TestValidateJWT_WrongSecret(t *testing.T) {
	parsedUUID, err := uuid.Parse(userId)
	if err != nil {
		t.Fatalf("failed to parse test UUID: %v", err)
	}

	token, err := MakeJWT(parsedUUID, tokenSecret, expiresIn)

	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if _, err := ValidateJWT(token, "wrong secret"); err == nil {
		t.Fatal("expected error for wrong secret, got nil")
	}
}

func TestValidateJWT_Expired(t *testing.T) {
	parsedUUID, err := uuid.Parse(userId)
	if err != nil {
		t.Fatalf("failed to parse test UUID: %v", err)
	}

	token, err := MakeJWT(parsedUUID, tokenSecret, -1*time.Hour)
	if err != nil {
		t.Fatalf("expected no error, got %v", err)
	}

	if _, err := ValidateJWT(token, tokenSecret); err == nil {
		t.Fatal("expected error for expired token, got nil")
	}
}
