package user
import "github.com/gorilla/mux"
import "net/http"

type Handler struct {
	// Add any dependencies here, e.g., database connection
}

func NewHandler() *Handler {
	return &Handler{
		// Initialize dependencies here
	}
}

func (h *Handler) RegisterRoutes(router *mux.Router) {
	router.HandleFunc("/login", h.handleLogin).Methods("GET")
}

func (h *Handler) handleLogin(w http.ResponseWriter, r *http.Request) {

}