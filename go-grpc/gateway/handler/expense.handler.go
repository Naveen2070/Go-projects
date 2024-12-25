package handler

import (
	"encoding/json"
	"net/http"

	expensepb "github.com/Naveen2070/Go-projects/go-grpc/common/api"
)

type Handler struct {
	client expensepb.ExpenseServiceClient
}

func NewHandler(client expensepb.ExpenseServiceClient) *Handler {
	return &Handler{client: client}
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/expenses", h.handleExpenses)
}

func (h *Handler) handleExpenses(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodPost:
		h.createExpense(w, r)
	case http.MethodGet:
		h.getExpenses(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *Handler) createExpense(w http.ResponseWriter, r *http.Request) {
	var req expensepb.CreateExpenseRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	resp, err := h.client.CreateExpense(r.Context(), &req)
	if err != nil {
		http.Error(w, "Failed to create expense: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *Handler) getExpenses(w http.ResponseWriter, r *http.Request) {
	req := &expensepb.GetExpensesRequest{}

	resp, err := h.client.GetExpenses(r.Context(), req)
	if err != nil {
		http.Error(w, "Failed to fetch expenses: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
