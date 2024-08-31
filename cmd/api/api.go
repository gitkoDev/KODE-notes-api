package api

import (
	"net/http"
	"time"
)

// type APIServer struct {
// 	addr string
// 	db   *sql.DB
// }

// func NewAPIServer(addr string, db *sql.DB) *APIServer {
// 	return &APIServer{addr, db}
// }

// func (s *APIServer) Run() error {
// 	router := chi.NewRouter()

// 	storage := service.NewStorage(s.db)
// 	handler := handler.NewHandler(storage)
// 	handler.RegisterRoutes(router)

// 	if err := http.ListenAndServe(s.addr, router); err != nil {
// 		return err
// 	}

// 	return nil
// }

type APIServer struct {
	httpServer *http.Server
}

func (s *APIServer) Run(port string, handler http.Handler) error {
	s.httpServer = &http.Server{
		Addr:           ":" + port,
		Handler:        handler,
		MaxHeaderBytes: 1 << 20, // 1 MB
		ReadTimeout:    10 * time.Second,
		WriteTimeout:   10 * time.Second,
	}

	if err := s.httpServer.ListenAndServe(); err != nil {
		return err
	}

	return nil
}
