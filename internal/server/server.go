package server

import (
	"net/http"

	"github.com/andreashanson/golang-awake/internal/handlers"
	"github.com/go-chi/chi"
)

type Server struct {
	Addr    string
	Handler http.Handler
}

func NewServer(p string, ih *handlers.InfluencerHandler) *Server {
	r := chi.NewRouter()
	r.Get("/status", http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.WriteHeader(200)
	}))
	r.Get("/influencers", http.HandlerFunc(ih.GetAll))
	r.Post("/influencers", http.HandlerFunc(ih.Create))
	r.Get("/influencers/{id}", http.HandlerFunc(ih.GetByID))

	return &Server{
		Addr:    ":" + p,
		Handler: r,
	}

}

func (s *Server) Run() error {
	srv := &http.Server{
		Addr:    s.Addr,
		Handler: s.Handler,
	}
	return srv.ListenAndServe()
}
