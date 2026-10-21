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
	// Récupérer la valeur du cache Redis
	data, err := config.RedisClient.Get(ctx, key).Bytes() // Récupérer la valeur en tant que tableau d'octets
	if err == redis.Nil {
		return false, nil
	} // La clé n'existe pas dans Redis
	if err != nil {
		return false, fmt.Errorf("erreur lors de la lecture du cache: %w", err)
	} // Une autre erreur s'est produite lors de la lecture du cache
	if err := json.Unmarshal(data, destination); err != nil { // Décoder la valeur JSON dans la destination
		return false, fmt.Errorf("erreur lors de la conversion du cache JSON: %w", err) // Une erreur s'est produite lors de la conversion du JSON en Go
	}

	// Si tout s'est bien passé, retourner true pour indiquer que la valeur a été trouvée et décodée avec succès
	return true, nil
}

// SetJSON encode value en JSON et l'enregistre dans Redis.
func SetJSON(ctx context.Context, key string, value any, expiration time.Duration) error {
	// Convertir la valeur en JSON
	data, err := json.Marshal(value)
	if err != nil {
		// Si une erreur se produit lors de la conversion en JSON, retourner une erreur
		return fmt.Errorf("erreur lors de la conversion en JSON: %w", err)
	} // Enregistrer la valeur JSON dans Redis avec une expiration
	if err := config.RedisClient.Set(ctx, key, data, expiration).Err(); err != nil {
		return fmt.Errorf("erreur lors de l'enregistrement du cache: %w", err)
	} // Si tout s'est bien passé, retourner nil pour indiquer que l'enregistrement a réussi
	return nil
}
