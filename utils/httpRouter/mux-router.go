package httpRouter

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi"
)

type muxRouter struct {
	port string
}

var (
	muxDispatcher = chi.NewRouter()
)

func NewMuxRouter() Router {
	return &muxRouter{}
}

func (*muxRouter) POST(uri string, f http.HandlerFunc) {
	muxDispatcher.Post(uri, f)
}

func (*muxRouter) GET(uri string, f http.HandlerFunc) {
	muxDispatcher.Get(uri, f)
}

func (*muxRouter) PUT(uri string, f http.HandlerFunc) {
	muxDispatcher.Put(uri, f)
}

func (*muxRouter) PATCH(uri string, f http.HandlerFunc) {
	muxDispatcher.Patch(uri, f)
}

func (*muxRouter) DELETE(uri string, f http.HandlerFunc) {
	muxDispatcher.Delete(uri, f)
}

func (m *muxRouter) SERVE(port string) {
	m.port = port

	fmt.Println("Online on http://localhost" + port)
	log.Fatal(http.ListenAndServe(port, muxDispatcher))
}

func (m *muxRouter) GetPortExposed() string {
	return m.port
}
