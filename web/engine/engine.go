package e

import (
	"fmt"
	"net/http"
)

// HandleFunc defines the return type in HTTP ServeHTTP
type HandleFunc func(http.ResponseWriter, *http.Request)

// Engine defines the core of the framework
// Implementation of ServeHTTP method, this engine will take over all
// the request from native HTTP package
type Engine struct {
	router map[string]HandleFunc
}

// New gets a plain engine with no configuration 
func New() *Engine {
	return &Engine{router: make(map[string]HandleFunc)}
}

// addRoute is used for adding request patterns and methods to the engine
// TODO: this can be chained
func (e *Engine) addRoute(method, pattern string, handler HandleFunc) {
	key := method + "-" + pattern
	e.router[key] = handler
}

// GET adds GET request to the engine
func (e *Engine) GET(pattern string, handler HandleFunc) {
	e.addRoute("GET", pattern, handler)
}

// POST adds POST request to the engine
func (e *Engine) POST(pattern string, handler HandleFunc) {
	e.addRoute("POST", pattern, handler)
}

// PUT adds PUT request to the engine
func (e *Engine) PUT(pattern string, handler HandleFunc) {
	e.addRoute("PUT", pattern, handler)
} 

// DELETE adds DELETE request to the engine
func (e *Engine) DELETE(pattern string, handler HandleFunc) {
	e.addRoute("PUT", pattern, handler)
} 

// ServeHTTP overrides the function in HTTP interface
// This takes over the control of interceptor in HTTP package
func (e *Engine) ServeHTTP(w http.ResponseWriter, req *http.Request) {
	key := req.Method + "-" + req.URL.Path
	if handler, ok := e.router[key]; ok {
		handler(w, req)
	} else {
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", req.URL)
	}
}

// Run executes the engine with the mapped routes and methods
func (e *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, e)
}

