package handler

import (
	"encoding/json"
	"net/http"

	userspb "github.com/Naveen2070/Go-projects/e-shop/common-service/users"
)

type Handler struct {
	client userspb.UserServiceClient
}

func NewHandler(client userspb.UserServiceClient) *Handler {
	return &Handler{client: client}
}

func (h *Handler) RegisterRoutes(mux *http.ServeMux) {
	mux.HandleFunc("/users", h.handleUsers)
}

func (h *Handler) handleUsers(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		h.getUsers(w, r)
	case http.MethodPost:
		h.createUser(w, r)
	default:
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func (h *Handler) createUser(w http.ResponseWriter, r *http.Request) {
	var req userspb.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	resp, err := h.client.CreateUser(r.Context(), &req)
	if err != nil {
		http.Error(w, "Failed to create expense: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

func (h *Handler) getUsers(w http.ResponseWriter, r *http.Request) {
	req := &userspb.GetUserByIdRequest{}

	resp, err := h.client.GetUserById(r.Context(), req)
	if err != nil {
		http.Error(w, "Failed to fetch expenses: "+err.Error(), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
