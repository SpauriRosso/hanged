package server

import (
	"hanged/src/shared"
	"net/http"
	"path/filepath"
	"time"
)

type Server struct {
	mux         *http.ServeMux
	staticDir   string
	templateDir string
}

func NewServer() *Server {
	return &Server{
		mux:         http.NewServeMux(),
		staticDir:   filepath.Join(shared.BasePath, "assets/static"),
		templateDir: filepath.Join(shared.BasePath, "assets/static/template"),
	}
}

func (s *Server) Start(addr string) error {
	s.Routes()
	server := &http.Server{
		Addr:              addr,
		Handler:           s.mux,
		ReadTimeout:       15 * time.Second,
		ReadHeaderTimeout: 5 * time.Second,
		WriteTimeout:      15 * time.Second,
		IdleTimeout:       60 * time.Second,
		MaxHeaderBytes:    1 << 20, // 1MB
	}
	return server.ListenAndServe()
}
