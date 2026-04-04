package handlers

import (
	"finboss/internal/models"
	"net/http"
)

type InvestmentHandler struct {
	repo InvestmentRepository
}

func NewInvestmentHandler(repo InvestmentRepository) *InvestmentHandler {
	return &InvestmentHandler{repo: repo}
}

func (h *InvestmentHandler) List(w http.ResponseWriter, r *http.Request) {
	investments, err := h.repo.FindAll()
	if err != nil {
		respondDBError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, investments)
}

func (h *InvestmentHandler) Get(w http.ResponseWriter, r *http.Request) {
	id, err := pathID(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid id")
		return
	}
	investment, err := h.repo.FindByID(id)
	if err != nil {
		respondDBError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, investment)
}

func (h *InvestmentHandler) Create(w http.ResponseWriter, r *http.Request) {
	var investment models.Investment
	if err := readJSON(r, &investment); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}
	if !investment.Type.IsValid() {
		writeError(w, http.StatusBadRequest, "invalid type")
		return
	}
	if err := h.repo.Create(&investment); err != nil {
		respondDBError(w, err)
		return
	}
	writeJSON(w, http.StatusCreated, investment)
}

func (h *InvestmentHandler) Update(w http.ResponseWriter, r *http.Request) {
	id, err := pathID(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid id")
		return
	}
	existing, err := h.repo.FindByID(id)
	if err != nil {
		respondDBError(w, err)
		return
	}
	if err := readJSON(r, &existing); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}
	if !existing.Type.IsValid() {
		writeError(w, http.StatusBadRequest, "invalid type")
		return
	}
	existing.ID = id
	if err := h.repo.Update(existing); err != nil {
		respondDBError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, existing)
}

func (h *InvestmentHandler) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := pathID(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid id")
		return
	}
	if err := h.repo.Delete(id); err != nil {
		respondDBError(w, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}
