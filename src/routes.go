package main

import (
	"encoding/json"
	"io"
	"net/http"
	"os"
	"os/exec"
	"path/filepath"
	"strconv"
	"time"

	"github.com/satori/go.uuid"
)

func mountRoutes() {
	router.Handle("/", baseChain.ThenFunc(convertHandler)).Methods(http.MethodPost, http.MethodPut)
	router.Handle("/health", baseChain.ThenFunc(healthCheckHandler)).Methods(http.MethodGet)
}

func healthCheckHandler(w http.ResponseWriter, r *http.Request) {
	d := map[string]interface{}{
		"uptime": time.Since(start).String(),
	}
	json.NewEncoder(w).Encode(&d)
}

func convertHandler(w http.ResponseWriter, r *http.Request) {
	tempDir, err := ensureTempDir()
	if err != nil {
		panic(err)
	}

	artifacts := []string{}

	defer func() {
		for _, filePath := range artifacts {
			if err := os.Remove(filePath); err != nil {
				logger.Errorw("failed to remove temporary file",
					"error", err,
					"path", filePath,
				)
			}
		}

		if recoverErr := recover(); recoverErr != nil {
			w.WriteHeader(http.StatusInternalServerError)

			if err := json.NewEncoder(w).Encode(map[string]interface{}{
				"error": recoverErr,
			}); err != nil {
				logger.Errorw("failed to respond with error",
					"originalError", recoverErr,
					"error", err,
				)
			}
		}
	}()

	tempSourceFileName := uuid.NewV4().String()
	tempSourceFilePath := filepath.Join(tempDir, tempSourceFileName)
	tempSourceFile, err := os.Create(tempSourceFilePath)
	if err != nil {
		panic(err)
	}

	artifacts = append(artifacts, tempSourceFilePath)

	receivedBytes, err := io.Copy(tempSourceFile, r.Body)
	if err != nil {
		panic(err)
	}

	if err := tempSourceFile.Close(); err != nil {
		panic(err)
	}

	tempTargetFileName := uuid.NewV4().String()
	tempTargetFilePath := filepath.Join(tempDir, tempTargetFileName)

	q := strconv.Itoa(config.Quality)

	cmd := exec.Command(config.ExecutablePath,
		"-quality", q,
		"-outfile", tempTargetFilePath,
		tempSourceFilePath,
	)

	if err := cmd.Run(); err != nil {
		panic(err)
	}

	artifacts = append(artifacts, tempTargetFilePath)

	stats, err := os.Stat(tempTargetFilePath)
	if err != nil {
		panic(err)
	}

	targetBytes := stats.Size()

	if targetBytes >= receivedBytes {
		logger.Infow("should send original file",
			"receivedBytes", receivedBytes,
			"targetBytes", targetBytes,
		)
		w.WriteHeader(http.StatusNoContent)
		return
	}

	tempTargetFile, err := os.Open(tempTargetFilePath)
	if err != nil {
		panic(err)
	}

	if _, err := io.Copy(w, tempTargetFile); err != nil {
		panic(err)
	}
}
