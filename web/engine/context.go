package engine

import (
	"encoding/json"
	"net/http"
	"strings"
)

// H simplifies the map structure
// Inspired by Gin
type H map[string]interface{}

// Context defines a request context
// This includes the lifecycle of a request from beginning to finish
type Context struct {
	Writer     http.ResponseWriter
	Request    *http.Request
	Path       string
	Method     string
	StatusCode int
}

// NewContext gives a new context
func NewContext(w http.ResponseWriter, req *http.Request) *Context {
	return &Context{
		Writer:  w,
		Request: req,
		Path:    req.URL.Path,
		Method:  req.Method,
	}
}

// PostForm gets the form data from the request
func (c *Context) PostForm(key string) string {
	return c.Request.FormValue(key)
}

// Query gets the request of the given key
func (c *Context) Query(key string) string {
	return c.Request.URL.Query().Get(key)
}

// Status sets the return HTTP status code
func (c *Context) Status(code int) {
	c.StatusCode = code
	c.Writer.WriteHeader(code)
}

// SetHeader sets the header of the request
func (c *Context) SetHeader(key string, value string) {
	c.Writer.Header().Set(key, value)
}

// String sets the response content with the given string
func (c *Context) String(code int, format string, values ...interface{}) {
	c.SetHeader("Content-Type", "text/plain")
	c.Status(code)
	sb := strings.Builder{}
	sb.WriteString(format)
	for _, v := range values {
		sb.WriteString(v.(string))
	}
	c.Writer.Write([]byte(sb.String()))
}

// JSON sets the response content with the given JSON object
func (c *Context) JSON(code int, obj interface{}) {
	c.SetHeader("Content-Type", "application/json")
	c.Status(code)
	encoder := json.NewEncoder(c.Writer)
	if err := encoder.Encode(obj); err != nil {
		http.Error(c.Writer, err.Error(), 500)
	}
}

// Data sets the response content with raw type of data
func (c *Context) Data(code int, data []byte) {
	c.Status(code)
	c.Writer.Write(data)
}

// HTML sets the response content with given HTML template
func (c *Context) HTML(code int, html string) {
	c.SetHeader("Content-Type", "text/html")
	c.Status(code)
	c.Writer.Write([]byte(html))
}
