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
	chiDispatcher = chi.NewRouter()
)

func NewChiRouter() Router {
	return &chiRouter{}
}

func (*chiRouter) POST(uri string, f http.HandlerFunc) {
	chiDispatcher.Post(uri, f)
}

func (*chiRouter) GET(uri string, f http.HandlerFunc) {
	chiDispatcher.Get(uri, f)
}

func (*chiRouter) PUT(uri string, f http.HandlerFunc) {
	chiDispatcher.Put(uri, f)
}

func (*chiRouter) PATCH(uri string, f http.HandlerFunc) {
	chiDispatcher.Patch(uri, f)
}

func (*chiRouter) DELETE(uri string, f http.HandlerFunc) {
	chiDispatcher.Delete(uri, f)
}

func (m *chiRouter) SERVE(port string) {
	fmt.Println("Online on http://localhost" + port)
	log.Fatal(http.ListenAndServe(port, chiDispatcher))
}
