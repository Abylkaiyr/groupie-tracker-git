package app

import (
	"fmt"
	"html/template"
	"net/http"
)

func (app *AppServer) ParseAndExecuteTemp(w http.ResponseWriter, file string,
	templateData interface{},
) {
	filePath := "ui/templates/" + file
	tmpl, err := template.ParseFiles(filePath)
	if err != nil {
		app.Errors(w, http.StatusInternalServerError,
			fmt.Errorf("parsing error %v html template: %w", file, err))
		return
	}

	err = tmpl.ExecuteTemplate(w, file, templateData)
	if err != nil {
		app.Errors(w, http.StatusInternalServerError,
			fmt.Errorf("executing error %v html template: %w", file, err))
		return
	}
}
