package handler

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/gomes800/password-manager/model"
	"github.com/gomes800/password-manager/repository"
)

type CredentialHandler struct {
	repo *repository.CredentialRepository
}

func NewCredentialHandler(repo *repository.CredentialRepository) *CredentialHandler {
	return &CredentialHandler{repo: repo}
}

func (h *CredentialHandler) Save(w http.ResponseWriter, r *http.Request) {
	var c model.Credential
	if err := json.NewDecoder(r.Body).Decode(&c); err != nil {
		http.Error(w, "Invalid request payload", http.StatusBadRequest)
		return
	}

	if err := h.repo.Save(r.Context(), &c); err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(c)
}

func (h *CredentialHandler) GetByID(w http.ResponseWriter, r *http.Request) {
	id := r.PathValue("id")

	credential, err := h.repo.GetByID(r.Context(), id)
	if err != nil {
		if err == sql.ErrNoRows {
			http.Error(w, "Credential not found", http.StatusNotFound)
			return
		}
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(credential)
}
