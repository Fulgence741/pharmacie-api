package services

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"pharmacie-api/models"

	"github.com/redis/go-redis/v9"
)

type AuditService struct {
	Redis *redis.Client
}

func NewAuditService(client *redis.Client) *AuditService {
	return &AuditService{
		Redis: client,
	}
}

func (s *AuditService) Record(
	ctx context.Context,
	audit models.AuditLog,
) error {
	if s.Redis == nil {
		return fmt.Errorf("client Redis non initialisé")
	}

	if audit.UserID == 0 {
		return fmt.Errorf("UserID est obligatoire")
	}

	if audit.Action == "" {
		return fmt.Errorf("Action est obligatoire")
	}

	if audit.Table == "" {
		return fmt.Errorf("Table est obligatoire")
	}

	if audit.Timestamp.IsZero() {
		audit.Timestamp = time.Now().UTC()
	}

	data, err := json.Marshal(audit)
	if err != nil {
		return fmt.Errorf("impossible de convertir l'audit en JSON : %w", err)
	}

	_, err = s.Redis.XAdd(ctx, &redis.XAddArgs{
		Stream: "audit:logs",
		Values: map[string]interface{}{
			"data": string(data),
		},
	}).Result()

	if err != nil {
		return fmt.Errorf("impossible d'enregistrer l'audit dans Redis : %w", err)
	}

	return nil
}
