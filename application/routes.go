package main

import (
	"github.com/julienschmidt/httprouter"
	"github.com/justinas/alice"
	"net/http"
	"path"
)

func (app *application) routes() http.Handler {
	md := alice.New(app.traceIds, app.requestLogger)

	mx := httprouter.New()
	mx.Handler("GET", "/socket.io", http.HandlerFunc(app.socketConn))
	mx.Handler("GET", "/", http.HandlerFunc(app.home))
	mx.ServeFiles("/static/*filepath", http.Dir(path.Join(app.Conf.ExecutablePath, "ui/static")))

	return md.Then(mx)
}
