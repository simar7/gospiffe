package main

import (
	"net/http"

	"github.com/go-chi/chi"
	"github.com/go-chi/chi/middleware"
)

type Server struct {
	Addr string
}

func (s *Server) NewRouter() *chi.Mux {
	r := chi.NewRouter()

	r.Use(middleware.RequestID)
	r.Use(middleware.RealIP)
	r.Use(middleware.Logger)
	r.Use(middleware.Recoverer)

	r.Get("/x509", FetchX509SVID)

	return r
}

func FetchJWTSVID(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte("FetchJWTSVID"))
}
func ValidateJWTSVID(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte("ValidateJWTSVID"))
}

func FetchJWTBundles(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte("FetchJWTBundles"))
}

func FetchX509SVID(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusNotImplemented)
	w.Write([]byte("FetchX509SVID"))
}

func main() {
	s := &Server{Addr: ":9991"}
	r := s.NewRouter()
	http.ListenAndServe(s.Addr, r)
}
