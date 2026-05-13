package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	fileserver := http.FileServer(http.Dir("."))

	mux.Handle("/app/", http.StripPrefix("/app", fileserver))

	mux.HandleFunc("/healthz", health)

	server := &http.Server{
		Addr:    ":8080",
		Handler: middlewareLog(mux),
	}

	fmt.Println("Server running on :8080")
	server.ListenAndServe()

}

func health(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "text/plain; charset=utf-8")

	w.WriteHeader(200)

	_, err := w.Write([]byte("OK"))

	if err != nil {
		fmt.Printf("Something went wrong writing bytes %w", err)
	}
}

func middlewareLog(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		log.Printf("%s %s", r.Method, r.URL.Path)

		next.ServeHTTP(w, r)
	})
}
