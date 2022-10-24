package main

import "github.com/gorilla/mux"

func NewMuxProvider() *mux.Router {
	return mux.NewRouter()
}
