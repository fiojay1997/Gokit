package engine

import (
	"net/http"
)

// HandleFunc defines the return type in HTTP ServeHTTP
type HandleFunc func(*Context)

// Engine defines the core of the framework
// Implementation of ServeHTTP method, this engine will take over all
// the request from native HTTP package
type Engine struct {
	router *router
}

// New gets a plain engine with no configuration
func New() *Engine {
	return &Engine{router: newRouter()}
}

// addRoute is used for adding request patterns and methods to the engine
// TODO: this can be chained
func (e *Engine) addRoute(method, pattern string, handler HandleFunc) {
	e.router.addRoute(method, pattern, handler)
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
	c := NewContext(w, req)
	e.router.getRoute(c)
}

// Run executes the engine with the mapped routes and methods
func (e *Engine) Run(addr string) (err error) {
	return http.ListenAndServe(addr, e)
}
