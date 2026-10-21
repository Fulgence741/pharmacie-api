package logging

import (
	"io"
	"log"
	"os"
	"path/filepath"
)

// InitLogger écrit les logs dans le fichier configuré et sur la sortie standard (stdout).

//	Elle crée le répertoire du fichier de log si nécessaire et configure le logger pour inclure

// les informations de date, d'heure et de fichier source dans les messages de log.
func InitLogger(logFile string) (*os.File, error) {
	// Création du répertoire du fichier de log si nécessaire
	if err := os.MkdirAll(filepath.Dir(logFile), 0755); err != nil {
		// Création du répertoire parent du fichier de log avec les permissions 0755
		return nil, err
	}

	// Ouverture du fichier de log en mode append, création s'il n'existe pas, avec les permissions 0644
	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	// Configuration du logger pour écrire à la fois dans le fichier de log et sur la sortie standard (stdout)
	log.SetOutput(io.MultiWriter(os.Stdout, file)) // Utilisation de io.MultiWriter pour écrire dans plusieurs destinations
	log.SetFlags(log.LstdFlags | log.Lshortfile)   // Configuration des flags du logger pour inclure la date, l'heure et le fichier source dans les messages de log
	return file, nil
}
