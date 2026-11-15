package main

import (
	"bytes"
	"fmt"
	"net/http"
	"runtime/debug"
)

// render looks up a page (e.g. "menu.html") in the template cache and runs
// its "base" layout. It renders into a buffer first, so a template error
// sends a clean 500 instead of half a page.
func (app *App) render(w http.ResponseWriter, status int, page string, data *templateData) {
	ts, ok := app.templateCache[page]
	if !ok {
		app.serverError(w, fmt.Errorf("the template %s does not exist", page))
		return
	}

	var buf bytes.Buffer
	if err := ts.ExecuteTemplate(&buf, "base", data); err != nil {
		app.serverError(w, err)
		return
	}

	w.WriteHeader(status)
	buf.WriteTo(w)
}

func (app *App) serverError(w http.ResponseWriter, err error) {
	trace := fmt.Sprintf("%s\n%s", err.Error(), debug.Stack())
	app.errorLog.Println(trace)
	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *App) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *App) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}
