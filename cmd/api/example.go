package main

import (
	"fmt"
	"net/http"

	"github.com/ozkansen/go-http-template/pkg/api"

	"github.com/go-chi/chi/v5"
)

var _ api.HttpAPI = (*ExampleHttpAPI)(nil)

type ExampleHttpAPI struct {
	api.DefaultHttpAPI
}

func (e *ExampleHttpAPI) ProductListHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	_, err := fmt.Fprint(w, `["a", "b", "c"]`)
	if err != nil {
		e.Logger().Error().Err(err).Send()
		return
	}
	e.Logger().Info().Msg("success")
}

func (e *ExampleHttpAPI) Router() *chi.Mux {
	router := e.DefaultHttpAPI.Router()
	router.Get("/", e.ProductListHandler)
	return router
}

func NewExampleAPI() *ExampleHttpAPI {
	e := &ExampleHttpAPI{}
	e.SetName("example")
	return e
}
