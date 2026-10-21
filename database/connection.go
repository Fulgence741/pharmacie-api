package database

import (
	"context"
	"fmt"
	"net"
	"net/url"
	"time"

	"pharmacie-api/config"

	"github.com/jackc/pgx/v5/pgxpool"
)

var DB *pgxpool.Pool

// ConnexionDB initialise et vérifie le pool PostgreSQL avant le démarrage HTTP.
func ConnexionDB(cfg *config.Config) error {
	connectionURL := url.URL{
		Scheme: "postgres",
		User:   url.UserPassword(cfg.DBUser, cfg.DBPassword),
		Host:   net.JoinHostPort(cfg.DBHost, cfg.DBPort),
		Path:   cfg.DBName,
		RawQuery: url.Values{
			"sslmode": []string{cfg.DBSSLMode},
		}.Encode(),
	}

	poolConfig, err := pgxpool.ParseConfig(connectionURL.String())
	if err != nil {
		return fmt.Errorf("configuration PostgreSQL invalide: %w", err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		return fmt.Errorf("création du pool PostgreSQL: %w", err)
	}
	if err := pool.Ping(ctx); err != nil {
		pool.Close()
		return fmt.Errorf("connexion à PostgreSQL: %w", err)
	}

	DB = pool
	return nil
}

// Close ferme le pool PostgreSQL initialisé au démarrage.
func Close() {
	if DB != nil {
		DB.Close()
		DB = nil
	}
}
