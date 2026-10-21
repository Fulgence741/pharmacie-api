package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"pharmacie-api/config"

	"github.com/redis/go-redis/v9"
)

// GetJSON récupère une valeur Redis et la décode depuis JSON.
func GetJSON(ctx context.Context, key string, destination any) (bool, error) {
	data, err := config.RedisClient.Get(ctx, key).Bytes()
	if err == redis.Nil {
		return false, nil
	}
	if err != nil {
		return false, fmt.Errorf("erreur lors de la lecture du cache: %w", err)
	}
	if err := json.Unmarshal(data, destination); err != nil {
		return false, fmt.Errorf("erreur lors de la conversion du cache JSON: %w", err)
	}
	return true, nil
}

// SetJSON encode value en JSON et l'enregistre dans Redis.
func SetJSON(ctx context.Context, key string, value any, expiration time.Duration) error {
	data, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("erreur lors de la conversion en JSON: %w", err)
	}
	if err := config.RedisClient.Set(ctx, key, data, expiration).Err(); err != nil {
		return fmt.Errorf("erreur lors de l'enregistrement du cache: %w", err)
	}
	return nil
}
