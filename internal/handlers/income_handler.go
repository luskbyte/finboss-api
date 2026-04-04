package handlers

import (
	"finboss/internal/models"
	"net/http"
)

type IncomeHandler struct {
	repo IncomeRepository
}

func NewIncomeHandler(repo IncomeRepository) *IncomeHandler {
	return &IncomeHandler{repo: repo}
}

func (h *IncomeHandler) List(w http.ResponseWriter, r *http.Request) {
	incomes, err := h.repo.FindAll()
	if err != nil {
		respondDBError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, incomes)
}

func (h *IncomeHandler) Get(w http.ResponseWriter, r *http.Request) {
	id, err := pathID(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid id")
		return
	}
	income, err := h.repo.FindByID(id)
	if err != nil {
		respondDBError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, income)
}

func (h *IncomeHandler) Create(w http.ResponseWriter, r *http.Request) {
	var income models.Income
	if err := readJSON(r, &income); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}
	if !income.Category.IsValid() {
		writeError(w, http.StatusBadRequest, "invalid category")
		return
	}
	if err := h.repo.Create(&income); err != nil {
		respondDBError(w, err)
		return
	}
	writeJSON(w, http.StatusCreated, income)
}

func (h *IncomeHandler) Update(w http.ResponseWriter, r *http.Request) {
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
	if err := readJSON(r, existing); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}
	if !existing.Category.IsValid() {
		writeError(w, http.StatusBadRequest, "invalid category")
		return
	}
	existing.ID = id
	if err := h.repo.Update(existing); err != nil {
		respondDBError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, existing)
}

func (h *IncomeHandler) Delete(w http.ResponseWriter, r *http.Request) {
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
