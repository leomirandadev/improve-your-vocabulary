package httpRouter

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

type chiRouter struct{}

var (
	chiDispatcher = chi.NewRouter()
)

func NewChiRouter() Router {
	return &chiRouter{}
}

func (*chiRouter) POST(uri string, f HandlerFunc) {
	chiDispatcher.Post(uri, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		f(newChiContext(w, r))
	})
}

func (*chiRouter) GET(uri string, f HandlerFunc) {
	chiDispatcher.Get(uri, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		f(newChiContext(w, r))
	})
}

func (*chiRouter) PUT(uri string, f HandlerFunc) {
	chiDispatcher.Put(uri, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		f(newChiContext(w, r))
	})
}

func (*chiRouter) PATCH(uri string, f HandlerFunc) {
	chiDispatcher.Patch(uri, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		f(newChiContext(w, r))
	})
}

func (*chiRouter) DELETE(uri string, f HandlerFunc) {
	chiDispatcher.Delete(uri, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		f(newChiContext(w, r))
	})
}

func (m *chiRouter) SERVE(port string) {
	fmt.Println("Online on http://localhost" + port)
	log.Fatal(http.ListenAndServe(port, chiDispatcher))
}

func (m *chiRouter) ParseHandler(h http.HandlerFunc) HandlerFunc {
	return func(c Context) {
		h(c.GetResponseWriter(), c.GetRequestReader())
	}
}

type chiContext struct {
	w http.ResponseWriter
	r *http.Request
}

func newChiContext(w http.ResponseWriter, r *http.Request) Context {
	return chiContext{w, r}
}

func (c chiContext) Context() context.Context {
	return c.r.Context()
}

func (c chiContext) JSON(status int, data interface{}) {
	c.w.WriteHeader(status)
	json.NewEncoder(c.w).Encode(data)
}

func (c chiContext) GetParam(param string) string {
	return chi.URLParam(c.r, "id")
}

func (c chiContext) GetFromHeader(param string) string {
	return c.r.Header.Get(param)
}

func (c chiContext) Decode(v interface{}) {
	json.NewDecoder(c.r.Body).Decode(v)
}

func (c chiContext) Headers() http.Header {
	return c.r.Header
}

func (c chiContext) GetResponseWriter() http.ResponseWriter {
	return c.w
}

func (c chiContext) GetRequestReader() *http.Request {
	return c.r
}
