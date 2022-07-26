package main

import (
	"net/http"
)

func (app *application) socketConn(w http.ResponseWriter, r *http.Request) {

	_, err := w.Write([]byte("Hello from 'setConn' handler"))
	if err != nil {
		app.Logger.Err(err, r.Context())
	}
}

func (app *application) home(w http.ResponseWriter, r *http.Request) {

	app.render(w, r, "home.page.tpl", &templateData{})
}
