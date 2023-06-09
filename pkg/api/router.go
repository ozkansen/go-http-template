package api

import (
	"net/http"

	"github.com/go-chi/chi/v5"
)

type Router interface {
	Add(p string, a HttpAPI)
	Use(m ...func(handler http.Handler) http.Handler)
	Router() (*chi.Mux, error)
}

var _ Router = (*DefaultRouter)(nil)

type DefaultRouter struct {
	store  map[string]HttpAPI
	router *chi.Mux
}

func (d *DefaultRouter) Add(p string, a HttpAPI) {
	d.store[p] = a
}

func (d *DefaultRouter) Use(m ...func(handler http.Handler) http.Handler) {
	d.router.Use(m...)
}

func (d *DefaultRouter) Router() (*chi.Mux, error) {
	for p, a := range d.store {
		if err := a.Configure(); err != nil {
			return nil, err
		}
		d.router.Mount(p, a.Router())
	}
	return d.router, nil
}

func NewRouter() *DefaultRouter {
	return &DefaultRouter{
		store:  make(map[string]HttpAPI, 100),
		router: chi.NewRouter(),
	}
}
