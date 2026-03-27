package middleware

import (
	"testing"

	"github.com/nashirabbash/backend-pfd/internal/config"
)

func TestGenerateAndValidateToken(t *testing.T) {
	config.AppConfig = &config.Config{
		JWTSecret:     "test-secret",
		JWTExpiration: "24",
	}

	token, err := GenerateToken(1, "user@example.com")
	if err != nil {
		t.Fatalf("GenerateToken() error = %v", err)
	}

	claims, err := ValidateToken(token)
	if err != nil {
		t.Fatalf("ValidateToken() error = %v", err)
	}

	if claims.UserID != 1 {
		t.Fatalf("expected user_id 1, got %d", claims.UserID)
	}

	if claims.Email != "user@example.com" {
		t.Fatalf("expected email user@example.com, got %s", claims.Email)
	}
}

func TestGenerateTokenFallbackForInvalidExpiration(t *testing.T) {
	config.AppConfig = &config.Config{
		JWTSecret:     "test-secret",
		JWTExpiration: "abc",
	}

	_, err := GenerateToken(1, "user@example.com")
	if err != nil {
		t.Fatalf("GenerateToken() should fallback to default expiration, got error = %v", err)
	}
}

func TestExtractToken(t *testing.T) {
	token, err := ExtractToken("Bearer abc.def.ghi")
	if err != nil {
		t.Fatalf("ExtractToken() error = %v", err)
	}

	if token != "abc.def.ghi" {
		t.Fatalf("expected token abc.def.ghi, got %s", token)
	}
}
