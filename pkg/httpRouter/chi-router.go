package httpRouter

import (
	"context"
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

type chiRouter struct {
	router *chi.Mux
}

func NewChiRouter() Router {
	return &chiRouter{
		router: chi.NewRouter(),
	}
}

func (r *chiRouter) POST(uri string, f HandlerFunc) {
	r.router.Post(uri, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		f(newChiContext(w, r))
	})
}

func (r *chiRouter) GET(uri string, f HandlerFunc) {
	r.router.Get(uri, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		f(newChiContext(w, r))
	})
}

func (r *chiRouter) PUT(uri string, f HandlerFunc) {
	r.router.Put(uri, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		f(newChiContext(w, r))
	})
}

func (r *chiRouter) PATCH(uri string, f HandlerFunc) {
	r.router.Patch(uri, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		f(newChiContext(w, r))
	})
}

func (r *chiRouter) DELETE(uri string, f HandlerFunc) {
	r.router.Delete(uri, func(w http.ResponseWriter, r *http.Request) {
		w.Header().Set("Content-Type", "application/json")
		f(newChiContext(w, r))
	})
}

func (r *chiRouter) SERVE(port string) {
	fmt.Println("Online on http://localhost" + port)
	log.Fatal(http.ListenAndServe(port, r.router))
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
	return chi.URLParam(c.r, param)
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
