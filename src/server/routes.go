package server

import "net/http"

func (s *Server) Routes() {
	//	Serve static files
	s.mux.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir(s.staticDir))))
	//	Routes
	s.mux.HandleFunc("/", s.RootHandler)
	s.mux.HandleFunc("/home", s.HomeHandler)
}
