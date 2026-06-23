package api

import (
	"database/sql"
	"github.com/gorilla/mux"
	"net/http"
	"github.com/dipeshpaneru/Go_Ecommerce/service/user"
	"log"
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



	userStore := user.NewStore(s.db)
	userHandler := user.NewHandler(userStore)
	userHandler.RegisterRoutes(subrouter)

	log.Printf("listening on %s", s.addr)

	return http.ListenAndServe(s.addr, router)


}