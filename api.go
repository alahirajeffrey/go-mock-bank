package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func WriteJson (w http.ResponseWriter, status int, v any) error {
	w.WriteHeader(status)
	w.Header().Set("Content-Type", "application/json")
	return json.NewEncoder(w).Encode(v)
}

type apiFunc func (http.ResponseWriter, *http.Request) error

type ApiError struct {
	Error string
}

func makeHTTPHandlerFunc (f apiFunc) http.HandlerFunc {
	return func (w http.ResponseWriter, r *http.Request)  {
		if err := f(w, r); err != nil{
			// handle error
			WriteJson(w, http.StatusBadRequest, ApiError{Error: err.Error()})
		}
	}
}
type APIServer struct {
	listenAddr string
}

func NewAPIServer(listenAddr string) *APIServer {
	return &APIServer{
		listenAddr: listenAddr,
	}
}

func (s *APIServer) Run () {
	router := mux.NewRouter()

	log.Println("Server running on port : ", s.listenAddr)
	router.HandleFunc("/account", makeHTTPHandlerFunc(s.handleAccount))

	http.ListenAndServe(s.listenAddr, router)
}

func (s *APIServer) handleAccount(w http.ResponseWriter, r *http.Request) error {
	if r.Method == "GET" {s.handleGetAccount(w, r)}

	if r.Method == "POST" {s.handleCreateAccount(w, r)}

	if r.Method == "DELETE" {s.handleDeleteAccount(w, r)}

	return fmt.Errorf("method not allowed %s", r.Method)	
 }

func (s *APIServer) handleGetAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleCreateAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleDeleteAccount(w http.ResponseWriter, r *http.Request) error {
	return nil
}

func (s *APIServer) handleTransfer(w http.ResponseWriter, r *http.Request) error {
	return nil
}