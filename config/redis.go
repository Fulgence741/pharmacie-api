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
	client := redis.NewClient(&redis.Options{
		Addr: address,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := client.Ping(ctx).Err(); err != nil {
		_ = client.Close()
		return fmt.Errorf("connexion à Redis (%s): %w", address, err)
	}

	RedisClient = client
	return nil
}

// CloseRedis ferme le client Redis initialisé au démarrage.
func CloseRedis() error {
	if RedisClient == nil {
		return nil
	}
	if err := RedisClient.Close(); err != nil {
		return fmt.Errorf("fermeture de la connexion Redis: %w", err)
	}
	RedisClient = nil
	return nil
}
