package main

import (
	"encoding/json"
	"net/http"
)

func mountRoutes() {
	router.Handle("/test", baseChain.ThenFunc(testHandler)).Methods(http.MethodGet)
}

func testHandler(w http.ResponseWriter, r *http.Request) {
	d := map[string]interface{}{"success": true}
	json.NewEncoder(w).Encode(&d)
}
