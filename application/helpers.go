package main

import (
	"bytes"
	"fmt"
	"net/http"
)

func (app *application) serverError(w http.ResponseWriter, r *http.Request, err error) {
	app.Logger.Err(err, r.Context())

	http.Error(w, http.StatusText(http.StatusInternalServerError), http.StatusInternalServerError)
}

func (app *application) clientError(w http.ResponseWriter, status int) {
	http.Error(w, http.StatusText(status), status)
}

func (app *application) notFound(w http.ResponseWriter) {
	app.clientError(w, http.StatusNotFound)
}

func (app *application) render(w http.ResponseWriter, r *http.Request, name string, td *templateData) {
	ts, ok := app.templateCache[name]
	if !ok {
		err := fmt.Errorf("template '%s' not found", name)
		app.serverError(w, r, err)
		return
	}

	buff := new(bytes.Buffer)
	err := ts.Execute(buff, td)
	if err != nil {
		app.serverError(w, r, err)
		return
	}

	_, err = buff.WriteTo(w)
	if err != nil {
		app.serverError(w, r, err)
	}
}
