package config

import "testing"

// TestLoadConfigUsesDefaults vérifie que LoadConfig()
//
//	utilise les valeurs par défaut lorsque certaines variables d'environnement ne sont pas définies.
func TestLoadConfigUsesDefaults(t *testing.T) {
	// Définir les variables d'environnement nécessaires pour le test,
	//  mais laisser certaines vides pour tester les valeurs par défaut
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

	// Charger la configuration
	cfg, err := LoadConfig()
	if err != nil {
		t.Fatalf("LoadConfig() returned an error: %v", err)
	}
	// Vérifier que les valeurs par défaut sont utilisées pour les variables d'environnement non définies

	if cfg.DBPort != "5432" || cfg.DBSSLMode != "disable" { // Vérifier les valeurs par défaut de PostgreSQL
		t.Errorf("unexpected PostgreSQL defaults: port=%q sslmode=%q", cfg.DBPort, cfg.DBSSLMode)
	} // Vérifier les valeurs par défaut de l'application
	if cfg.ServerPort != "8080" || cfg.RedisAddr != "127.0.0.1:6379" || cfg.LogFile != "logs/app.log" {
		t.Errorf("unexpected application defaults: port=%q Redis=%q log=%q", cfg.ServerPort, cfg.RedisAddr, cfg.LogFile)
	}
}

// TestLoadConfigRequiresJWTSecret vérifie que LoadConfig() rejette une configuration sans JWT_SECRET.
func TestLoadConfigRequiresJWTSecret(t *testing.T) {
	// Définir les variables d'environnement nécessaires,
	// mais laisser JWT_SECRET vide
	t.Setenv("DB_HOST", "localhost")
	t.Setenv("DB_USER", "test-user")
	t.Setenv("DB_NAME", "test-db")
	t.Setenv("JWT_SECRET", "")
	// Charger la configuration et vérifier qu'une erreur est renvoyée
	if _, err := LoadConfig(); err == nil {
		t.Fatal("LoadConfig() should reject a missing JWT_SECRET")
	}
}
