package main

import (
	"net"
	"net/http"
	"strconv"
	"time"

	"github.com/gorilla/mux"
	"github.com/justinas/alice"
	"go.uber.org/zap"
)

var (
	start     = time.Now()
	config    = Config{}
	router    = mux.NewRouter().StrictSlash(true)
	baseChain = alice.New(recoverHandler)
	logger    *zap.SugaredLogger
)

func main() {
	mustConfigure()
	mustInitLogger()
	mountRoutes()

	host := config.Host
	port := strconv.Itoa(config.Port)

	server := http.Server{
		Addr:    net.JoinHostPort(host, port),
		Handler: router,
	}

	logger.Infow("server started",
		"host", host,
		"port", port,
		"duration", time.Since(start),
	)

	if err := server.ListenAndServe(); err != nil {
		panic(err)
	}
}
