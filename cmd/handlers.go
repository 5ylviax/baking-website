package main

import (
	"net/http"
	"os"
	"path/filepath"
)

// The menu is the home page, so "/" shows it (the Let's Go home handler).
// "/" also catches every unknown path, so anything else is a 404.
func (app *App) handleIndex(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		app.notFound(w)
		return
	}
	app.handleMenu(w, r)
}

func (app *App) handleMenu(w http.ResponseWriter, r *http.Request) {
	menu := make([]menuCategory, len(app.menu))
	for i, c := range app.menu {
		items := make([]menuItem, len(c.Items))
		for j, item := range c.Items {
			item.Photos = existingPhotos(item.Images)
			items[j] = item
		}
		c.Items = items
		menu[i] = c
	}

	app.render(w, "menu", menu)
}

// existingPhotos returns the names that have a file in static/img.
func existingPhotos(names []string) []string {
	var found []string
	for _, name := range names {
		if _, err := os.Stat(filepath.Join("static", "img", name)); err == nil {
			found = append(found, name)
		}
	}
	return found
}

func (app *App) handleAbout(w http.ResponseWriter, r *http.Request) {
	app.render(w, "about", nil)
}

func (app *App) handleContact(w http.ResponseWriter, r *http.Request) {
	app.render(w, "contact", nil)
}
