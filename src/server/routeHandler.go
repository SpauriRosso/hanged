package server

import (
	"html/template"
	"log"
	"net/http"
	"path/filepath"
)

func (s *Server) RootHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/" {
		http.Redirect(w, r, "/home", http.StatusSeeOther)
	}
}

func (s *Server) HomeHandler(w http.ResponseWriter, r *http.Request) {
	tmp, err := template.ParseFiles(filepath.Join(s.templateDir, "index.html"))
	if err != nil {
		log.Println(err)
		return
	}
	err = tmp.Execute(w, r)
	if err != nil {
		log.Println(err)
		return
	}
}
