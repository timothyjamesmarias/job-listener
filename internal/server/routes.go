package server

import (
	"encoding/json"
	"job-listener/internal/database"
	_ "job-listener/internal/database"
	"job-listener/internal/database/models"
	"log"
	"net/http"
	"strconv"
)

func (s *Server) RegisterRoutes() http.Handler {
	mux := http.NewServeMux()

	// Register routes
	mux.HandleFunc("/health", s.healthHandler)

	// Page routes

	mux.HandleFunc("/", s.dashboardHandler)

	// API routes

	// App model resource routes
	mux.HandleFunc("GET /api/v1/apps", MakeHTTPHandleFunc(s.getAppsHandler))
	mux.HandleFunc("POST /api/v1/apps", MakeHTTPHandleFunc(s.createAppHandler))
	mux.HandleFunc("GET /api/v1/apps/{id}", MakeHTTPHandleFunc(s.getAppByIDHandler))
	mux.HandleFunc("PUT /api/v1/apps/{id}", MakeHTTPHandleFunc(s.updateAppHandler))
	mux.HandleFunc("DELETE /api/v1/apps/{id}", MakeHTTPHandleFunc(s.destroyAppHandler))

	// Wrap the mux with CORS middleware
	return s.corsMiddleware(mux)
}

func (s *Server) corsMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		// Set CORS headers
		w.Header().Set("Access-Control-Allow-Origin", "*") // Replace "*" with specific origins if needed
		w.Header().Set("Access-Control-Allow-Methods", "GET, POST, PUT, DELETE, OPTIONS, PATCH")
		w.Header().Set("Access-Control-Allow-Headers", "Accept, Authorization, Content-Type, X-CSRF-Token")
		w.Header().Set("Access-Control-Allow-Credentials", "false") // Set to "true" if credentials are required

		// Handle preflight OPTIONS requests
		if r.Method == http.MethodOptions {
			w.WriteHeader(http.StatusNoContent)
			return
		}

		// Proceed with the next handler
		next.ServeHTTP(w, r)
	})
}

func (s *Server) dashboardHandler(w http.ResponseWriter, r *http.Request) {
	resp := map[string]string{"message": "Hello World"}
	jsonResp, err := json.Marshal(resp)
	if err != nil {
		http.Error(w, "Failed to marshal response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(jsonResp); err != nil {
		log.Printf("Failed to write response: %v", err)
	}
}

func (s *Server) healthHandler(w http.ResponseWriter, r *http.Request) {
	resp, err := json.Marshal(database.Health(s.db))
	if err != nil {
		http.Error(w, "Failed to marshal health check response", http.StatusInternalServerError)
		return
	}
	w.Header().Set("Content-Type", "application/json")
	if _, err := w.Write(resp); err != nil {
		log.Printf("Failed to write response: %v", err)
	}
}

func (s *Server) getAppsHandler(w http.ResponseWriter, r *http.Request) error {
	apps, err := models.GetAllApps(s.db)
	if err != nil {
		return WriteJSON(w, http.StatusNotFound, apiError{Error: "Not found"})
	}
	return WriteJSON(w, http.StatusOK, apps)
}

func (s *Server) createAppHandler(w http.ResponseWriter, r *http.Request) error {
	return WriteJSON(w, http.StatusOK, nil)
}

func (s *Server) getAppByIDHandler(w http.ResponseWriter, r *http.Request) error {
	id, err := strconv.Atoi(r.PathValue("id"))
	if err != nil {
		return WriteJSON(w, http.StatusNotFound, apiError{Error: "App not found"})
	}

	app, err := models.GetAppByID(s.db, id)

	if err != nil {
		return WriteJSON(w, http.StatusNotFound, apiError{Error: "App not found"})
	}

	return WriteJSON(w, http.StatusOK, app)
}

func (s *Server) updateAppHandler(w http.ResponseWriter, r *http.Request) error {
	return WriteJSON(w, http.StatusOK, nil)
}

func (s *Server) destroyAppHandler(w http.ResponseWriter, r *http.Request) error {
	return WriteJSON(w, http.StatusOK, nil)
}
