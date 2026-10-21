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
	// Construction de l'URL de connexion PostgreSQL
	connectionURL := url.URL{
		Scheme: "postgres",                                   // le schéma de l'URL pour PostgreSQL
		User:   url.UserPassword(cfg.DBUser, cfg.DBPassword), // les informations d'identification pour la connexion à PostgreSQL
		Host:   net.JoinHostPort(cfg.DBHost, cfg.DBPort),     // l'adresse de l'hôte et le port pour la connexion à PostgreSQL
		Path:   cfg.DBName,                                   // le nom de la base de données PostgreSQL
		RawQuery: url.Values{ // les paramètres de requête pour la connexion à PostgreSQL
			"sslmode": []string{cfg.DBSSLMode}, // le mode SSL pour la connexion à PostgreSQL
		}.Encode(), // encodage des paramètres de requête en chaîne de caractères
	}

	// Analyse de l'URL de connexion PostgreSQL pour obtenir la configuration du pool
	poolConfig, err := pgxpool.ParseConfig(connectionURL.String()) // analyse de l'URL de connexion PostgreSQL pour obtenir la configuration du pool
	if err != nil {
		return fmt.Errorf("configuration PostgreSQL invalide: %w", err)
	}

	// Définition des paramètres du pool PostgreSQL
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second) // création d'un contexte avec un délai d'attente de
	//  5 secondes pour la connexion à PostgreSQL
	defer cancel() // annulation du contexte après la connexion à PostgreSQL

	// Création du pool PostgreSQL avec la configuration obtenue
	pool, err := pgxpool.NewWithConfig(ctx, poolConfig)
	if err != nil {
		return fmt.Errorf("création du pool PostgreSQL: %w", err)
	}
	// Vérification de la connexion à PostgreSQL en effectuant un ping
	if err := pool.Ping(ctx); err != nil {
		pool.Close() // fermeture du pool PostgreSQL en cas d'erreur de ping
		return fmt.Errorf("connexion à PostgreSQL: %w", err)
	}

	// Attribution du pool PostgreSQL à la variable globale DB pour une utilisation ultérieure
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
