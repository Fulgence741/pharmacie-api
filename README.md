# Pharmacie API

API HTTP en Go pour gérer les pharmacies, les gardes, les utilisateurs et les liens entre pharmacies et gardes. Les routes sont déclarées dans `routes`, et la documentation est disponible sur `/swagger/`.

## Démarrer le projet

Prérequis : Go 1.26 ou plus récent et une base PostgreSQL accessible.

Créer un fichier `.env` à la racine avec les paramètres nécessaires :

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=votre_utilisateur
DB_PASSWORD=votre_mot_de_passe
DB_NAME=votre_base
DB_SSLMODE=disable
JWT_SECRET=une_cle_secrete
```

Puis lancer l'API :

```bash
go run .
```

Le serveur écoute actuellement sur le port `8080`.
## Routes principales

Toutes les routes métier utilisent la méthode `POST`.

 Ressource et Routes 

 Pharmacies  `/pharmacie`, `/pharmacie/list`, `/pharmacie/{id_pharmacie}/get`, `/pharmacie/{id_pharmacie}/put`, `/pharmacie/{id_pharmacie}/delete` 
 Gardes  `/garde`, `/garde/list`, `/garde/{id_garde}/get`, `/garde/{id_garde}/put`, `/garde/{id_garde}/delete` |
 Pharmacies de garde  `/pharmacie-garde`, `/pharmacie-garde/list`, `/pharmacie-garde/{id}/delete` 
 Utilisateurs  `/user`, `/user/list`, `/user/{id}/get`, `/user/{id}/delete`, `/user/{id}/role/patch` 
 Authentification  `/login` 

Les routes protégées attendent un jeton JWT dans l'en-tête `Authorization`, avec le schéma `Bearer` suivi du jeton. Les listes acceptent les paramètres de pagination `page` et `limit`, ainsi que leurs filtres de recherche.

## Organisation du code

- `handlers/` reçoit les requêtes HTTP, utilise les DTO et transmet le contexte de la requête.
- `dto/` définit séparément les formats JSON reçus (`requests.go`) et renvoyés (`responses.go`).
- `services/` applique les validations et les règles métier.
- `repositories/` exécute les requêtes PostgreSQL.
- `models/` décrit les données manipulées par les services et les repositories.
- `middleware/` regroupe l'authentification, les rôles, les limites de requêtes et les logs.
- `errors/` centralise les erreurs métier/HTTP et leur format de réponse JSON.
- `database/` charge la configuration, ouvre la connexion et gère les migrations.

## Contexte et erreurs

Pour une requête entrante, les handlers transmettent `request.Context()` aux services, puis aux repositories. Le contexte porte  l'annulation du client et les échéances : les opérations SQL doivent réutiliser ce contexte et ne pas le remplacer par `context.Background()`. Un contexte avec délai indépendant est réservé aux opérations d'initialisation, comme la connexion à PostgreSQL.

Les erreurs HTTP sont écrites par le point commun dans `errors/` au format `{"error":"..."}`. Les erreurs de validation, d'authentification, d'autorisation et de ressource absente conservent leur statut HTTP ; les erreurs internes sont consignées dans les logs et ne révèlent pas les détails de la base de données au client.

Les DTO d'entrée empêchent de lier directement le JSON reçu aux modèles de base de données. Les DTO de sortie permettent de contrôler les champs exposés ; par exemple, la réponse utilisateur ne contient pas le mot de passe.

## Vérification

```bash
go test ./...
```