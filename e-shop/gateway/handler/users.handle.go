package handler

import (
	"encoding/json"
	"net/http"

	userspb "github.com/Naveen2070/Go-projects/e-shop/common-service/users"
	"github.com/gorilla/mux"
)

type Handler struct {
	client userspb.UserServiceClient
}

// NewHandler initializes a new Handler instance.
func NewHandler(client userspb.UserServiceClient) *Handler {
	return &Handler{client: client}
}

// RegisterRoutes registers the API routes with the given router.
func (h *Handler) RegisterRoutes(router *mux.Router) {
	// Define routes
	router.HandleFunc("/users", h.createUser).Methods(http.MethodPost)
	router.HandleFunc("/users", h.getAllUsers).Methods(http.MethodGet)
	router.HandleFunc("/users/{user_id}", h.getUserByID).Methods(http.MethodGet)
}

// createUser handles the POST /users endpoint.
func (h *Handler) createUser(w http.ResponseWriter, r *http.Request) {
	var req userspb.CreateUserRequest
	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		http.Error(w, "Invalid request body", http.StatusBadRequest)
		return
	}

	// Call the gRPC service
	resp, err := h.client.CreateUser(r.Context(), &req)
	if err != nil {
		http.Error(w, "Failed to create user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Return the response as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// getAllUsers handles the GET /users endpoint.
func (h *Handler) getAllUsers(w http.ResponseWriter, r *http.Request) {
	// Call the gRPC service
	resp, err := h.client.GetUsers(r.Context(), &userspb.GetUsersRequest{})
	if err != nil {
		if resp == nil {
			http.Error(w, "No users found", http.StatusNotFound)
			return
		}
		http.Error(w, "Failed to get users: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Return the response as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}

// getUserByID handles the GET /users/{user_id} endpoint.
func (h *Handler) getUserByID(w http.ResponseWriter, r *http.Request) {
	// Extract user_id from the URL path
	vars := mux.Vars(r)
	userID, ok := vars["user_id"]
	if !ok {
		http.Error(w, "user_id parameter is missing", http.StatusBadRequest)
		return
	}

	req := &userspb.GetUserByIdRequest{
		UserId: userID,
	}

	// Call the gRPC service
	resp, err := h.client.GetUserById(r.Context(), req)
	if err != nil {
		http.Error(w, "Failed to fetch user: "+err.Error(), http.StatusInternalServerError)
		return
	}

	// Return the response as JSON
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(resp)
}
