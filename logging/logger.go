package logging

import (
	"fmt"
	"log"
	"os"
)

// InitLogger initialise le logger pour écrire les logs dans un fichier.
func InitLogger(logFile string) {
	// Ouvre le fichier de log en mode append, crée-le s'il n'existe pas, et définit les permissions à 0644
	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Fatalf("Erreur lors de l'ouverture du fichier de log: %v", err)
	}

	// Définit la sortie du logger vers le fichier ouvert
	log.SetOutput(file)

	// Définit le format du logger pour inclure la date, l'heure et le nom du fichier
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	// Affiche un message indiquant que le logger a été initialisé avec succès
	fmt.Printf("Logger initialisé. Les logs seront écrits dans %s\n", logFile)
}
