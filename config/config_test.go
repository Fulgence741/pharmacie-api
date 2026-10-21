package config

import "testing"

func TestLoadConfigUsesDefaults(t *testing.T) {
	t.Setenv("DB_HOST", "localhost")
	t.Setenv("DB_USER", "test-user")
	t.Setenv("DB_PASSWORD", "")
	t.Setenv("DB_NAME", "test-db")
	t.Setenv("JWT_SECRET", "test-secret")
	t.Setenv("DB_PORT", "")
	t.Setenv("DB_SSLMODE", "")
	t.Setenv("SERVER_PORT", "")
	t.Setenv("REDIS_ADDR", "")
	t.Setenv("LOG_FILE", "")

	cfg, err := LoadConfig()
	if err != nil {
		t.Fatalf("LoadConfig() returned an error: %v", err)
	}

	if cfg.DBPort != "5432" || cfg.DBSSLMode != "disable" {
		t.Errorf("unexpected PostgreSQL defaults: port=%q sslmode=%q", cfg.DBPort, cfg.DBSSLMode)
	}
	if cfg.ServerPort != "8080" || cfg.RedisAddr != "127.0.0.1:6379" || cfg.LogFile != "logs/app.log" {
		t.Errorf("unexpected application defaults: port=%q Redis=%q log=%q", cfg.ServerPort, cfg.RedisAddr, cfg.LogFile)
	}
}

func TestLoadConfigRequiresJWTSecret(t *testing.T) {
	t.Setenv("DB_HOST", "localhost")
	t.Setenv("DB_USER", "test-user")
	t.Setenv("DB_NAME", "test-db")
	t.Setenv("JWT_SECRET", "")

	if _, err := LoadConfig(); err == nil {
		t.Fatal("LoadConfig() should reject a missing JWT_SECRET")
	}
}
