package main

import "html/template"

// parseTemplates loads every template once at startup into a single set.
// ParseGlob doesn't search subfolders, so each folder gets its own glob.
// Every {{define "name"}} must be unique across all files, or a later
// file silently replaces an earlier one.
func parseTemplates() (*template.Template, error) {
	ts, err := template.ParseGlob("./templates/page/*.html")
	if err != nil {
		return nil, err
	}
	return ts.ParseGlob("./templates/partials/*.html")
}
