package main

import (
	"net/http"
	"os"
	"path/filepath"

	"panaderia.de.santiago.net/internal/services"
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
	menu := make([]services.MenuCategory, len(app.menu))
	for i, c := range app.menu {
		items := make([]services.MenuItem, len(c.Items))
		for j, item := range c.Items {
			item.Photos = existingPhotos(item.Images)
			items[j] = item
		}
		c.Items = items
		menu[i] = c
	}

	app.render(w, http.StatusOK, "menu.html", &templateData{Menu: menu})
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
	app.render(w, http.StatusOK, "about.html", &templateData{})
}

func (app *App) handleContact(w http.ResponseWriter, r *http.Request) {
	app.render(w, http.StatusOK, "contact.html", &templateData{})
}
