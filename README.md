# Pharmacie API

API HTTP en Go pour gérer les pharmacies, les gardes, les pharmacies de garde et les utilisateurs. La documentation Swagger est servie sur `/swagger/`.

## Prérequis

- Go 1.26.5 ou version ultérieure.
- PostgreSQL accessible depuis l'application.
- Redis accessible depuis l'application.

## Configuration

Copiez `.env.example` vers `.env`, puis renseignez les paramètres PostgreSQL et le secret JWT :

```env
DB_HOST=localhost
DB_PORT=5432
DB_USER=votre_utilisateur
DB_PASSWORD=votre_mot_de_passe
DB_NAME=votre_base
DB_SSLMODE=disable
JWT_SECRET=remplacez_par_un_secret_long_et_aleatoire
SERVER_PORT=8080
REDIS_ADDR=127.0.0.1:6379
LOG_FILE=logs/app.log
```

`DB_HOST`, `DB_USER`, `DB_NAME` et `JWT_SECRET` sont obligatoires. Les autres valeurs ont des valeurs par défaut. En dehors de Docker, l'application utilise Redis sur `127.0.0.1:6379`.

## Démarrage

En local, après avoir démarré PostgreSQL et Redis :

```bash
go run .
```

Le serveur écoute sur le port `8080` par défaut. `SERVER_PORT` permet de le modifier.

Avec Docker Compose, Redis est lancé automatiquement. PostgreSQL reste un service externe ; sous Docker Desktop, utilisez généralement `DB_HOST=host.docker.internal` pour joindre PostgreSQL installé sur l'hôte :

```bash
docker compose up --build
```

## Routes principales

Les routes métier utilisent actuellement `POST`.

Pharmacies : `/pharmacie`, `/pharmacie/list`, `/pharmacie/{id}/get`, `/pharmacie/{id}/put`, `/pharmacie/{id}/delete`

Gardes : `/garde`, `/garde/list`, `/garde/{id}/get`, `/garde/{id}/put`, `/garde/{id}/delete`

Pharmacies de garde : `/pharmacie-garde`, `/pharmacie-garde/list`, `/pharmacie-garde/{id}/delete`

Utilisateurs : `/user`, `/user/list`, `/user/{id}/get`, `/user/{id}/delete`, `/user/{id}/role/patch`

Authentification : `/login`

Les routes protégées attendent un jeton JWT dans l'en-tête `Authorization`, au format `Bearer <jeton>`. Les routes de liste prennent en charge la pagination (`page`, `limit`) et les filtres définis par les handlers.

## Organisation du code

auth : création et vérification des jetons JWT
cache : gestion du cache Redis
config : lecture de la configuration et connexion à Redis
database : connexion à PostgreSQL
docs : documentation Swagger
dto : formats des données reçues et renvoyées par l'API
errors : gestion des erreurs HTTP et métier
handlers : traitement des requêtes HTTP
logging : configuration des logs
middleware : authentification, rôles, limitation des requêtes et journalisation
models : structures des données métier
repositories : accès aux données PostgreSQL
routes : déclaration des routes de l'API
services : validation et règles métier
utils : fonctions partagées, notamment la pagination
main.go : démarrage de l'application et connexion aux services
server.go : écoute des requêtes et arrêt propre du serveur

Les handlers transmettent le contexte HTTP aux services puis aux repositories afin que l'annulation des requêtes soit respectée. Les erreurs de connexion PostgreSQL ou Redis empêchent le serveur de démarrer. À l'arrêt, le serveur se ferme proprement et libère ses connexions.

Les logs sont envoyés à la sortie standard et au fichier défini par `LOG_FILE` (par défaut `logs/app.log`). Les erreurs internes sont consignées côté serveur sans exposer les détails techniques au client.

## Tests

```bash
go test ./...
```
