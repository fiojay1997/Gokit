package engine

import (
	"log"
	"net/http"
)

// router defines handler function
type router struct {
	handlers map[string]HandleFunc
}

// newRouter gives a new router object
func newRouter() *router {
	return &router{handlers: make(map[string]HandleFunc)}
}

// addRoute adds routes with given pattern to the request
func (r *router) addRoute(method string, pattern string, handler HandleFunc) {
	log.Printf("Route %4s - %s", method, pattern)
	key := method + "-" + pattern
	r.handlers[key] = handler
}

// handle process the request
func (r *router) handle(c *Context) {
	key := c.Method + "-" + c.Path
	if handler, ok := r.handlers[key]; ok {
		handler(c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND: %s\n", c.Path)
	}
}
