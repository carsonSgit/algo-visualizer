package main

import (
	"log"
	"net/http"
	"os"
	"path/filepath"

	"github.com/devuser/algo-visualizer/algorithms"
	"github.com/gorilla/mux"
)

func serveStatic(dir string) http.Handler {
	return http.FileServer(http.Dir(dir))
}

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/api/sort", algorithms.HandleSort).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/generate", algorithms.HandleGenerateArray).Methods("GET", "OPTIONS")

	nextOut := filepath.Join(".", "frontend", "out")
	viteDist := filepath.Join(".", "frontend", "dist")

	if _, err := os.Stat(nextOut); err == nil {
		router.PathPrefix("/").Handler(serveStatic(nextOut))
	} else if _, err := os.Stat(viteDist); err == nil {
		router.PathPrefix("/").Handler(serveStatic(viteDist))
	} else {
		router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
			w.Header().Set("Content-Type", "text/html")
			w.WriteHeader(http.StatusOK)
			w.Write([]byte(`
				<!DOCTYPE html>
				<html>
					<head>
						<title>Algorithm Visualizer</title>
						<meta charset="UTF-8">
						<meta name="viewport" content="width=device-width, initial-scale=1.0">
					</head>
					<body>
						<div id="root"></div>
						<p>Build the frontend with: pnpm build</p>
					</body>
				</html>
			`))
		})
	}

	log.Println("Server starting on http://localhost:8080")
	if err := http.ListenAndServe(":8080", router); err != nil {
		log.Fatal(err)
	}
}
