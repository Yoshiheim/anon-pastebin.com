package handlers

import (
	"net/http"
	"text/template"
)

func The404Handler(w http.ResponseWriter, r *http.Request) {
	templ, err := template.ParseFiles("./templates/404.html")
	if err != nil {
		http.Error(w, "error with HTML", http.StatusNotAcceptable)
		return
	}
	err = templ.Execute(w, nil)
	if err != nil{
		http.Error(w, "Error with execute HTML template.", http.StatusNotAcceptable)
	}
}
