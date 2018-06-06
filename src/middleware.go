package main

import (
	"fmt"
	"net/http"
)

func recoverHandler(h http.Handler) http.Handler {
	return http.HandlerFunc(
		func(w http.ResponseWriter, r *http.Request) {
			defer func() {
				switch err := recover(); err {
				case nil:
					return
				default:
					switch err.(type) {
					default:
						w.WriteHeader(http.StatusInternalServerError)

						if config.Env == "development" {
							fmt.Fprintf(w, "%v", err)
						}

						logger.Errorw("recovered from unhandled error",
							"error", err,
						)
					}
				}
			}()
			h.ServeHTTP(w, r)
		},
	)
}
