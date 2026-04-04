package handlers

import (
	"errors"
	"finboss/internal/repositories"
	"log"
	"net/http"
)

func respondDBError(w http.ResponseWriter, err error) {
	if errors.Is(err, repositories.ErrNotFound) {
		writeError(w, http.StatusNotFound, "not found")
		return
	}
	log.Printf("internal error: %v", err)
	writeError(w, http.StatusInternalServerError, "internal server error")
}
