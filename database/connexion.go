package database

import (
	"context"
	"fmt"
	"time"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

func ConnexionDB(config *Config) error {

	// construction du dsn pour créer une chaîne de caractère utilisé pour la connexion
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		config.DBUser,
		config.DBPassword,
		config.DBHost,
		config.DBPort,
		config.DBName,
		config.DBSSLMode,
	)

	// S'assurer que le dsn a été bien formé
	poolConfig, err := pgxpool.ParseConfig(dsn)
	if err != nil {
		return fmt.Errorf("Dsn mal formé : %w", err)
	}

	//Créer un context pour les délais

	ctx, cancel := context.WithTimeout(
		context.Background(),
		5*time.Second,
	)

	defer cancel()

	//Création du pool de connexion à partir de la configuration précédente
	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		return fmt.Errorf("Erreur de creation du pool : %w", err)
	}

	//Vérifier que le pool pour communiquer avec postgreSQL
	err = pool.Ping(ctx)

	if err != nil {
		pool.Close()
		return fmt.Errorf("Impossible de joindre postgreSQL : %w", err)
	}

	// Mettre le pool dans ma variable globale
	DB = pool

	return nil

}
