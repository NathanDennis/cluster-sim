package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {

	handler := http.NewServeMux()

	const numEndpoints = 1

	for i := 0; i < numEndpoints; i++ {
		route := fmt.Sprintf("/cluster/%d", i)

		handler.HandleFunc(route, func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "application/json")

			json.NewEncoder(w).Encode(map[string]string{
				"routeName": route,
				"message":   route + " hit",
			})

			log.Printf("[%s] %s - %s %s", r.RemoteAddr, r.Method, r.Host, r.URL.Path)
		})
	}

	s := &http.Server{
		Addr:         ":1337",
		Handler:      handler,
		ReadTimeout:  10 * time.Second,
		WriteTimeout: 10 * time.Second,
	}

	fmt.Printf("Server listening on %s\n", s.Addr)
	s.ListenAndServe()
}
