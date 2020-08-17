package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func main() {

	handler := http.NewServeMux()

	for i := 0; i < 10; i++ {
		route := fmt.Sprintf("/cluster/%d", i)

		handler.HandleFunc(route, func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")

			json.NewEncoder(w).Encode(map[string]string{
				"routeName": route,
				"message":   route + " hit",
			})
		})
	}

	s := &http.Server{
		Addr:         ":1337",
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	s.ListenAndServe()
}
