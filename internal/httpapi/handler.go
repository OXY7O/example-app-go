package httpapi

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type healthResponse struct {
	Status string `json:"status"`
}

type exampleResponse struct {
	ID   int    `json:"id"`
	Name string `json:"name"`
}

type errorResponse struct {
	Error string `json:"error"`
}

func NewHandler() http.Handler {
	mux := http.NewServeMux()
	mux.HandleFunc("/health", health)
	mux.HandleFunc("/examples/", exampleByID)
	return mux
}

func health(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSON(w, http.StatusMethodNotAllowed, errorResponse{Error: "method not allowed"})
		return
	}
	writeJSON(w, http.StatusOK, healthResponse{Status: "ok"})
}

func exampleByID(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		writeJSON(w, http.StatusMethodNotAllowed, errorResponse{Error: "method not allowed"})
		return
	}
	value := strings.TrimPrefix(r.URL.Path, "/examples/")
	id, err := strconv.Atoi(value)
	if err != nil || id < 1 || strings.Contains(value, "/") {
		writeJSON(w, http.StatusBadRequest, errorResponse{Error: "invalid example id"})
		return
	}
	writeJSON(w, http.StatusOK, exampleResponse{ID: id, Name: "example-" + strconv.Itoa(id)})
}

func writeJSON(w http.ResponseWriter, status int, body any) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	_ = json.NewEncoder(w).Encode(body)
}
