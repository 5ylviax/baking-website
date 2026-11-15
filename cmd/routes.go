package main

import "net/http"

func (app *App) routes() *http.ServeMux {
	mux := http.NewServeMux()

	mux.HandleFunc("/", app.handleIndex)
	mux.HandleFunc("GET /menu", app.handleMenu)
	mux.HandleFunc("GET /about", app.handleAbout)
	mux.HandleFunc("GET /contact", app.handleContact)

	//CSS
	mux.Handle("/static/", http.StripPrefix("/static", http.FileServer(http.Dir("static"))))

	return mux
}
