package main

import (
	"crypto/tls"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	app := NewApp()
	if app == nil {
		log.Fatalln("Application not created")
	}

	tlsConfig := &tls.Config{
		CurvePreferences: []tls.CurveID{tls.X25519, tls.CurveP256},
	}

	addr := fmt.Sprintf("%s:%s", app.Conf.ListenAddr, app.Conf.ListenPort)

	srv := &http.Server{
		Addr: addr,
		//ErrorLog:     app.Logger, TODO: Привести к нужному контракту? Или просто не задавать?
		Handler:      app.routes(),
		TLSConfig:    tlsConfig,
		IdleTimeout:  time.Minute,
		ReadTimeout:  5 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	app.Logger.Info(fmt.Sprintf("Server listening in %s", addr), nil)
	err := srv.ListenAndServeTLS(app.Conf.CertFilePath, app.Conf.KeyFilePath)
	app.Logger.Err(err, nil)
}
