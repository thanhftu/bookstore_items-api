package app

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/thanhftu/bookstore_items-api/clients/elasticsearch"
)

var (
	router = mux.NewRouter()
)

// StartApplication for start app
func StartApplication() {
	elasticsearch.Init()
	mapURL()
	srv := &http.Server{
		Handler: router,
		Addr:    "127.0.0.1:8081",
		// Good practice: enforce timeouts for servers you create!
		WriteTimeout: 500 * time.Millisecond,
		ReadTimeout:  2 * time.Second,
		IdleTimeout:  60 * time.Second,
	}
	fmt.Println("app is running on port:8081")

	if err := srv.ListenAndServe(); err != nil {

		panic(err)
	}
}
