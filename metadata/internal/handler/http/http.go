package http

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"movieexample.com/metadata/internal/controller"
	"movieexample.com/metadata/internal/repository"
)

// Handler defines a movie metadata HTTP controller.
type Handler struct {
	ctrl *controller.Controller
}

// New creates a new movie metadata HTTP handler.
func New(ctrl *controller.Controller) *Handler {
	return &Handler{ctrl}
}

// GetMetadata handle GET /metadata requests.
func (h *Handler) GetMetadata(w http.ResponseWriter, r *http.Request) {
	id := r.FormValue("id")
	if id == "" {
		w.WriteHeader(http.StatusNotFound)
		return
	}
	ctx := r.Context()
	m, err := h.ctrl.Get(ctx, id)
	if err != nil && errors.Is(err, repository.ErrNotFound) {
		w.WriteHeader(http.StatusNotFound)
		return
	} else if err != nil {
		log.Printf("Repository get error: %v\n", err)
		w.WriteHeader(http.StatusInternalServerError)
		return
	}
	if err := json.NewEncoder(w).Encode(m); err != nil {
		log.Printf("Response encode error: %v\n", err)
	}
}
