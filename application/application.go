package main

import (
	"git.slygods.com/evoplay/wss-go/pkg/helper"
	"git.slygods.com/evoplay/wss-go/service/wsslogger"
	"html/template"
	"io"
	"os"
	"path"
)

const (
	keyRequestId    = wsslogger.ContextKey("request_id") // Это приведение типов, а не вызов функции
	keySpanId       = wsslogger.ContextKey("span_id")
	keyParentSpanId = wsslogger.ContextKey("parent_span_id")
)

type application struct {
	Logger        *wsslogger.Logger
	Conf          *config
	templateCache map[string]*template.Template
}

func NewApp() *application {
	app := new(application)

	logger := wsslogger.NewLogger(
		helper.Env("SERVICE_NAME", "WSS_V2"),
		[]io.Writer{os.Stdout},
		[]wsslogger.ContextKey{
			keyRequestId,
			keySpanId,
			keyParentSpanId,
		})

	conf, err := NewConfig()
	if err != nil {
		logger.Err(err, nil)
		return nil
	}

	tc, err := NewTemplateCache(path.Join(conf.ExecutablePath, "ui/html"))
	if err != nil {
		logger.Err(err, nil)
		return nil
	}

	app.Logger = logger
	app.Conf = conf
	app.templateCache = tc

	return app
}
