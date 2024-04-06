package server

import (
	"github.com/codegangsta/negroni"
	"github.com/diogosilvasantana/go-hexagonal/adapters/web/handler"
	"github.com/diogosilvasantana/go-hexagonal/application"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"os"
	"time"
)

type WebServer struct {
	Service application.ProductServiceInterface
}

func MakeNewWebServer() WebServer {
	return WebServer{}
}

func (w WebServer) Serve() {

	r := mux.NewRouter()
	n := negroni.New(
		negroni.NewLogger(),
	)

	handler.MakeProductHandler(r, n, w.Service)
	http.Handle("/", r)

	server := &http.Server{
		ReadHeaderTimeout: 10 * time.Second,
		WriteTimeout:      10 * time.Second,
		Addr:              ":9000",
		Handler:           nil,
		ErrorLog:          log.New(os.Stderr, "log: ", log.Lshortfile),
	}
	err := server.ListenAndServe()
	if err != nil {
		log.Fatal(err)
	}
}
