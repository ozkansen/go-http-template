package api

import (
	"go-http-template/pkg/logger"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"
)

type HttpAPI interface {
	// Name çalışan api yi isimlendirin
	Name() string
	SetName(n string)

	// Logger api ye tanımlanmış logger
	Logger() *zerolog.Logger
	SetLogger(l *zerolog.Logger)

	// Init tanımlamalar bittikten sonra çalıştırın
	Configure() error

	// Router api yazıldıktan sonra endpoint tanımlamaları burada yapılır
	Router() *chi.Mux
}

var _ HttpAPI = (*DefaultHttpAPI)(nil)

type DefaultHttpAPI struct {
	router *chi.Mux
	logger *zerolog.Logger
	name   string
}

func (d *DefaultHttpAPI) Name() string {
	return d.name
}

func (d *DefaultHttpAPI) SetName(n string) {
	d.name = n
}

func (d *DefaultHttpAPI) Logger() *zerolog.Logger {
	return d.logger
}

func (d *DefaultHttpAPI) SetLogger(l *zerolog.Logger) {
	d.logger = l
}

func (d *DefaultHttpAPI) Configure() error {
	if d.logger == nil {
		d.logger = logger.NewDefaultConsoleLogger()
	}
	l := d.logger.With().Str("api", d.Name()).Logger()
	d.logger = &l
	d.router = chi.NewRouter()
	return nil
}

func (d *DefaultHttpAPI) Router() *chi.Mux {
	return d.router
}
