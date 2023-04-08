package api

import (
	"go-http-template/pkg/logger"

	"github.com/go-chi/chi/v5"
	"github.com/rs/zerolog"
)

type HttpAPI interface {
	Name() string
	SetName(n string)
	Logger() *zerolog.Logger
	SetLogger(l *zerolog.Logger)
	Init() error
	Router() *chi.Mux
}

var _ HttpAPI = (*DefaultHttpAPI)(nil)

type DefaultHttpAPI struct {
	name   string
	router *chi.Mux
	logger *zerolog.Logger
}

func (d *DefaultHttpAPI) Name() string {
	return ""
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

func (d *DefaultHttpAPI) Init() error {
	if d.logger == nil {
		d.logger = logger.NewDefaultConsoleLogger()
		d.logger.Info().Str("name", d.Name()).Msg("set default logger")
	}
	d.router = chi.NewRouter()
	return nil
}

func (d *DefaultHttpAPI) Router() *chi.Mux {
	panic("implement me")
}
