package config

import "testing"

func TestGetDSNUsesDatabaseURLWhenAvailable(t *testing.T) {
	cfg := &Config{DatabaseURL: "postgres://user:pass@localhost:5432/db?sslmode=disable"}

	if got := cfg.GetDSN(); got != cfg.DatabaseURL {
		t.Fatalf("expected DATABASE_URL to be returned, got %s", got)
	}
}

func TestGetDSNFallbacksToLegacyFields(t *testing.T) {
	cfg := &Config{
		DBHost:     "localhost",
		DBPort:     "5432",
		DBUser:     "postgres",
		DBPassword: "postgres",
		DBName:     "pfd_db",
	}

	want := "host=localhost port=5432 user=postgres password=postgres dbname=pfd_db sslmode=disable"
	if got := cfg.GetDSN(); got != want {
		t.Fatalf("expected %s, got %s", want, got)
	}
}
