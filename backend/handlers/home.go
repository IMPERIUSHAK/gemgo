package handlers

import (
	"html/template"
	"net/http"
	"path/filepath"
)

func HandlerFunc(w http.ResponseWriter, r *http.Request) {

	homeTemplate, err := template.ParseFiles(filepath.Join("..", "templates", "index.html"))
	if err != nil {
		http.Error(w, "Error while loading page template", http.StatusInternalServerError)
	}

	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	if r.Method != http.MethodGet {
		http.Error(w, "Method not supported", http.StatusMethodNotAllowed)
	}

	if err := homeTemplate.Execute(w, ""); err != nil {
		http.Error(w, "Error while rendering home page", http.StatusInternalServerError)
	}
}
