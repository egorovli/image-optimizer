package main

import (
	"fmt"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

var (
	config = Config{}
	router = mux.NewRouter().StrictSlash(true)
	start  = time.Now()
)

func main() {
	server := http.Server{
		Addr:    ":8080",
		Handler: router,
	}

	fmt.Println("Server started")

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
