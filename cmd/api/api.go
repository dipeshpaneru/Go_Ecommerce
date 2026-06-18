package api

import (
	"database/sql",
	"github.com/gorilla/mux",
	"net/http"
)

type APIServer struct {
	addr string
	db *sql.DB
}

// Constructor for APIServer

func NewAPIServer(addr string, db *sql.DB) *APIServer {
	return &APIServer{
		addr: addr,
		db: db,
	}
}

// Run starts the API server
func (s *APIServer) Run() error {
	router := mux.NewRouter()
	subrouter := router.PathPrefix("/api/v1").Subrouter()

	log("listening on %s", s.addr)

	return http.ListenAndServe(s.addr, router)


}