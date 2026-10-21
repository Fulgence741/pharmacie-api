package config

import (
	"context"
	"fmt"
	"time"

	"github.com/redis/go-redis/v9"
)

var RedisClient *redis.Client

// InitRedis vérifie Redis avant de laisser l'application accepter des requêtes.
func InitRedis(address string) error {
	// Créer un nouveau client Redis avec l'adresse fournie
	client := redis.NewClient(&redis.Options{
		Addr: address,
	})

	// Créer un contexte avec un délai d'attente pour la vérification de la connexion
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	// Vérifier la connexion à Redis en envoyant un ping
	if err := client.Ping(ctx).Err(); err != nil {
		_ = client.Close() // Fermer le client Redis en cas d'erreur de connexion
		return fmt.Errorf("connexion à Redis (%s): %w", address, err)
	}

	// Si la connexion est réussie, assigner le client Redis global
	RedisClient = client
	return nil
}

// CloseRedis ferme le client Redis initialisé au démarrage.
func CloseRedis() error {
	if RedisClient == nil {
		return nil
	}
	// Fermer la connexion Redis et gérer les erreurs
	if err := RedisClient.Close(); err != nil {
		return fmt.Errorf("fermeture de la connexion Redis: %w", err)
	}
	// Réinitialiser le client Redis global à nil
	RedisClient = nil
	return nil
}
