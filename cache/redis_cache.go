package cache

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"pharmacie-api/config"

	"github.com/redis/go-redis/v9"
)

// GetJSON récupère une donnée dans Redis et la convertit en structure Go.
func GetJSON(
	ctx context.Context,
	key string,
	destination interface{},
) (bool, error) {

	// ÉTAPE 1 : rechercher la valeur associée à cette clé dans Redis.
	//
	// Bytes() permet de récupérer la donnée sous forme d'octets
	// ensuite convertir ces octets en structure Go.
	data, err := config.RedisClient.Get(ctx, key).Bytes()

	// ÉTAPE 2 : vérifier si la clé n'existe pas dans Redis
	// redis.Nil signifie  que la donnée n'est pas dans le cache
	if err == redis.Nil {
		return false, nil
	}

	// ÉTAPE 3 : vérifier si une véritable erreur Redis s'est produite

	if err != nil {
		return false, fmt.Errorf(
			"erreur lors de la lecture du cache : %w",
			err,
		)
	}

	// ÉTAPE 4 : convertir le JSON récupéré en structure Go
	// destination est un pointeur vers la variable qui recevra les données
	if err := json.Unmarshal(data, destination); err != nil {
		return false, fmt.Errorf(
			"erreur lors de la conversion du cache JSON : %w",
			err,
		)
	}

	// ÉTAPE 5 : la donnée existe et a été correctement convertie.

	//  retourner true pour indiquer que le cache contient le résultat
	return true, nil
}

// SetJSON enregistre une donnée dans Redis au format JSON
func SetJSON(
	ctx context.Context,
	key string,
	value interface{},
	expiration time.Duration,
) error {

	// ÉTAPE 1 : convertir la donnée Go en JSON

	data, err := json.Marshal(value)

	// ÉTAPE 2 : vérifier si la conversion en JSON a échoué.
	if err != nil {
		return fmt.Errorf(
			"erreur lors de la conversion en JSON : %w",
			err,
		)
	}

	// ÉTAPE 3 : enregistrer le JSON dans Redis
	// expiration représente la durée pendant laquelle la donnée sera conservée
	// Avec 5*time.Minute Redis supprimera la clé après 5 minutes.
	err = config.RedisClient.Set(
		ctx,
		key,
		data,
		expiration,
	).Err()

	// ÉTAPE 4 : vérifier si l'enregistrement dans Redis a échoué.
	if err != nil {
		return fmt.Errorf(
			"erreur lors de l'enregistrement du cache : %w",
			err,
		)
	}

	// ÉTAPE 5 : l'enregistrement s'est déroulé correctement.
	return nil
}

// Delete supprime une ou plusieurs clés du cache Redis.

func Delete(
	ctx context.Context,
	keys ...string,
) error {

	// ÉTAPE 1 : vérifier si aucune clé n'a été fournie.

	//    inutile d'interroger Redis si la liste est vide.
	if len(keys) == 0 {
		return nil
	}

	// ÉTAPE 2 : supprimer les clés dans Redis.
	// Del accepte plusieurs clés en un seul appel.
	err := config.RedisClient.Del(
		ctx,
		keys...,
	).Err()

	// ÉTAPE 3 : vérifier si Redis a rencontré une erreur.
	if err != nil {
		return fmt.Errorf(
			"erreur lors de la suppression du cache : %w",
			err,
		)
	}

	// ÉTAPE 4 : la suppression s'est déroulée correctement.
	return nil
}
