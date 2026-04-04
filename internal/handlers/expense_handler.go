package handlers

import (
	"finboss/internal/models"
	"net/http"
)

type ExpenseHandler struct {
	repo ExpenseRepository
}

func NewExpenseHandler(repo ExpenseRepository) *ExpenseHandler {
	return &ExpenseHandler{repo: repo}
}

func (h *ExpenseHandler) List(w http.ResponseWriter, r *http.Request) {
	expenses, err := h.repo.FindAll()
	if err != nil {
		respondDBError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, expenses)
}

func (h *ExpenseHandler) Get(w http.ResponseWriter, r *http.Request) {
	id, err := pathID(r)
	if err != nil {
		writeError(w, http.StatusBadRequest, "invalid id")
		return
	}
	expense, err := h.repo.FindByID(id)
	if err != nil {
		respondDBError(w, err)
		return
	}
	writeJSON(w, http.StatusOK, expense)
}

func (h *ExpenseHandler) Create(w http.ResponseWriter, r *http.Request) {
	var expense models.Expense
	if err := readJSON(r, &expense); err != nil {
		writeError(w, http.StatusBadRequest, "invalid request body")
		return
	}
	if !expense.Category.IsValid() {
		writeError(w, http.StatusBadRequest, "invalid category")
		return
	}
	if err := h.repo.Create(&expense); err != nil {
		respondDBError(w, err)
		return
	}
	writeJSON(w, http.StatusCreated, expense)
}

func (h *ExpenseHandler) Update(w http.ResponseWriter, r *http.Request) {
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

func (h *ExpenseHandler) Delete(w http.ResponseWriter, r *http.Request) {
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
