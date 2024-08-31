package handler

import (
	"net/http"

	"github.com/gitkoDev/KODE-test-task/internal/service"
	"github.com/go-chi/chi"
)

type Handler struct {
	service *service.Service
}

func NewHandler(service *service.Service) *Handler {
	return &Handler{service: service}
}

func (h *Handler) RegisterRoutes() http.Handler {
	router := chi.NewRouter()

	router.Get("/health", h.ping)

	router.Route("/auth", func(r chi.Router) {
		r.Post("/sign-up", h.signUp)
		r.Post("/sign-in", h.signIn)
	})

	router.With(h.IdentifyUser).Route("/api", func(r chi.Router) {
		r.Route("/v1", func(r chi.Router) {
			r.Get("/notes", h.getNotes)
			r.Post("/notes", h.addNote)
		})
	})

	return router

}
