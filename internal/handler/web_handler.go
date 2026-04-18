package handler

import (
	"net/http"
	"text/template"
)

func RenderLogin(w http.ResponseWriter, r *http.Request) {
	tmpl := template.Must(template.ParseFiles("web/template/login.html"))
	tmpl.Execute(w, nil)
}

func RenderDashboard(w http.ResponseWriter, r *http.Request){
	tmpl, err := template.ParseFiles("web/template/dashboard.html")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)
}

func RenderPOS(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("web/template/pos.html")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)
}

func RenderProducts(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("web/template/product.html")

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	tmpl.Execute(w, nil)
}