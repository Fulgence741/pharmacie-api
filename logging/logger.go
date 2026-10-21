package logging

import (
	"io"
	"log"
	"os"
	"path/filepath"
)

// InitLogger écrit les logs dans le fichier configuré et sur la sortie standard.
func InitLogger(logFile string) (*os.File, error) {
	if err := os.MkdirAll(filepath.Dir(logFile), 0755); err != nil {
		return nil, err
	}

	file, err := os.OpenFile(logFile, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		return nil, err
	}

	log.SetOutput(io.MultiWriter(os.Stdout, file))
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	return file, nil
}
