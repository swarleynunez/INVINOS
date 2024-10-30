package handlers

import (
	"html/template"
	"net/http"
)

func Index(w http.ResponseWriter, r *http.Request) {
	// Parse template
	tmpl, err := template.ParseFiles("app/static/index.html")
	if err != nil {
		http.Error(w, "Error:", http.StatusInternalServerError)
		return
	}

	// Render template
	err = tmpl.Execute(w, nil)
	if err != nil {
		http.Error(w, "Error:", http.StatusInternalServerError)
	}
}
