# Etape 1 : utiliser une image Go pour compiler l'application
FROM golang:1.26-alpine AS builder 

# Dossier de travail dans le conteneur 
WORKDIR /app 

# Copier les fichiers de dépedances 
COPY go.mod go.sum ./

# Télécharger les dépendances 
RUN go mod download 

# Copier tout le projet 
COPY . .

# Compiler l'API 
RUN go build -o pharmacie-api .

# Etape 2 : Créer une image légère pour exécuter l'API 
FROM alpine:latest 

WORKDIR /app

# Copier l'exécutable depuis l'étape précédente 
COPY --from=builder /app/pharmacie-api .

# L'API utilise le port 8080 
EXPOSE 8080 

# Demarre l'API
CMD ["./pharmacie-api"]
