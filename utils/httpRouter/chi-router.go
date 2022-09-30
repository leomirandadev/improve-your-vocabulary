package httpRouter

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

type chiRouter struct {
}

var (
	muxDispatcher = chi.NewRouter()
)

func NewMuxRouter() Router {
	return &chiRouter{}
}

func (*chiRouter) POST(uri string, f http.HandlerFunc) {
	muxDispatcher.Post(uri, f)
}

func (*chiRouter) GET(uri string, f http.HandlerFunc) {
	muxDispatcher.Get(uri, f)
}

func (*chiRouter) PUT(uri string, f http.HandlerFunc) {
	muxDispatcher.Put(uri, f)
}

func (*chiRouter) PATCH(uri string, f http.HandlerFunc) {
	muxDispatcher.Patch(uri, f)
}

func (*chiRouter) DELETE(uri string, f http.HandlerFunc) {
	muxDispatcher.Delete(uri, f)
}

func (m *chiRouter) SERVE(port string) {
	fmt.Println("Online on http://localhost" + port)
	log.Fatal(http.ListenAndServe(port, muxDispatcher))
}
