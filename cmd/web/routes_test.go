package main

import (
	"fmt"
	"testing"

	"github.com/go-chi/chi/v5"
	"github.com/iamlego/bookingBoost/internal/config"
)

func TestRoutes(t *testing.T) {
	var app config.AppConfig
	mux := routes(&app)

	switch v := mux.(type) {
	case *chi.Mux:
		//do nothing
	default:
		t.Errorf("%s", fmt.Sprintf("type is not http.Handler, instead its a %T", v))
	}
}
