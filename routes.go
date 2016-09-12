package main

import (
	"github.com/gorilla/mux"
	"net/http"
)

type Route struct {
	Method      string
	Path        string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = []Route{
	Route{
		Method:      "GET",
		Path:        "/",
		HandlerFunc: http.HandlerFunc(Index)},
	Route{
		Method:      "GET",
		Path:        "/list/{id}",
		HandlerFunc: http.HandlerFunc(GetTodo)},
	Route{
		Method:      "POST",
		Path:        "/insert",
		HandlerFunc: http.HandlerFunc(CreateTodo)},
	Route{
		Method:      "DELETE",
		Path:        "/list/{id}",
		HandlerFunc: http.HandlerFunc(DeleteTodo)},
	Route{
		Method:      "GET",
		Path:        "/list",
		HandlerFunc: http.HandlerFunc(ShowTodo)},
	Route{
		Method:      "PUT",
		Path:        "/list/{id}",
		HandlerFunc: http.HandlerFunc(UpdateTodo)},
	Route{
		Method:      "GET",
		Path:        "/save",
		HandlerFunc: http.HandlerFunc(SaveTodo)},
	Route{
		Method:      "GET",
		Path:        "/load",
		HandlerFunc: http.HandlerFunc(LoadTodo)},
}

func NewRouter() *mux.Router {
	// StrictSlash defines the trailing slash behavior for new routes. The initial value is false.
	// When true, if the route path is "/path/", accessing "/path" will redirect to the former and vice versa.
	// When false, if the route path is "/path", accessing "/path/" will not match this route and vice versa.

	router := mux.NewRouter().StrictSlash(true)

	for _, route := range routes {
		router.Methods(route.Method).
			Path(route.Path).
			HandlerFunc(route.HandlerFunc)
	}
	return router
}
