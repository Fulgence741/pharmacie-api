package config

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

// Config contient les paramètres requis au démarrage de l'application.
type Config struct {
	DBHost     string
	DBPort     string
	DBUser     string
	DBPassword string
	DBName     string
	DBSSLMode  string

	ServerPort string
	RedisAddr  string
	LogFile    string
}

// LoadConfig lit l'environnement et applique les valeurs par défaut documentées.
func LoadConfig() (*Config, error) {
	if err := godotenv.Load(); err != nil && !os.IsNotExist(err) {
		return nil, fmt.Errorf("chargement du fichier .env: %w", err)
	}

	// Vérifier que les variables d'environnement obligatoires sont définies
	required := []string{"DB_HOST", "DB_USER", "DB_NAME", "JWT_SECRET"}
	// Vérifier que toutes les variables d'environnement obligatoires sont définies
	for _, name := range required {
		// Vérifier si la variable d'environnement est vide
		if os.Getenv(name) == "" {
			return nil, fmt.Errorf("la variable d'environnement %s est obligatoire", name)
		}
	}

	// Retourner la configuration avec les valeurs par défaut pour les variables d'environnement non définies
	return &Config{
		DBHost:     os.Getenv("DB_HOST"),
		DBPort:     valueOrDefault("DB_PORT", "5432"),
		DBUser:     os.Getenv("DB_USER"),
		DBPassword: os.Getenv("DB_PASSWORD"),
		DBName:     os.Getenv("DB_NAME"),
		DBSSLMode:  valueOrDefault("DB_SSLMODE", "disable"),
		ServerPort: valueOrDefault("SERVER_PORT", "8080"),
		RedisAddr:  valueOrDefault("REDIS_ADDR", "127.0.0.1:6379"),
		LogFile:    valueOrDefault("LOG_FILE", "logs/app.log"),
	}, nil
}

// valueOrDefault retourne la valeur de la variable d'environnement
// si elle est définie, sinon elle retourne la valeur par défaut.

func valueOrDefault(name, defaultValue string) string {
	// Vérifier si la variable d'environnement est définie et non vide
	if value := os.Getenv(name); value != "" {
		return value
	}
	// Si la variable d'environnement n'est pas définie, retourner la valeur par défaut
	return defaultValue
}
