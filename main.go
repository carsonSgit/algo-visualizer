package main

import (
	"encoding/json"
	"log"
	"net/http"
	"os"
	"path/filepath"
	"strconv"
	"strings"

	"github.com/gorilla/mux"
)

// Step represents a single step in the sorting algorithm
type Step struct {
	Array      []int   `json:"array"`
	Comparing  []int   `json:"comparing"`
	Swapped    []int   `json:"swapped"`
	Sorted     []int   `json:"sorted"`
	StepNumber int     `json:"step_number"`
	Message    string  `json:"message"`
}

// SortResult contains all steps for a complete sort
type SortResult struct {
	Steps      []Step `json:"steps"`
	Algorithm  string `json:"algorithm"`
	Duration   string `json:"duration"`
	Comparisons int   `json:"comparisons"`
	Swaps      int    `json:"swaps"`
}

// BubbleSort performs bubble sort and returns all steps
func BubbleSort(arr []int) SortResult {
	result := SortResult{
		Steps:     make([]Step, 0),
		Algorithm: "Bubble Sort",
		Swaps:     0,
		Comparisons: 0,
	}

	arr = append([]int(nil), arr...) // Copy array
	n := len(arr)
	sorted := make([]bool, n)
	stepNum := 0

	// Initial state
	result.Steps = append(result.Steps, Step{
		Array:      append([]int(nil), arr...),
		Comparing:  []int{},
		Swapped:    []int{},
		Sorted:     []int{},
		StepNumber: stepNum,
		Message:    "Starting Bubble Sort",
	})
	stepNum++

	for i := 0; i < n-1; i++ {
		swapped := false

		for j := 0; j < n-i-1; j++ {
			result.Comparisons++

			// Comparing step
			result.Steps = append(result.Steps, Step{
				Array:      append([]int(nil), arr...),
				Comparing:  []int{j, j + 1},
				Swapped:    []int{},
				Sorted:     getSortedIndices(sorted),
				StepNumber: stepNum,
				Message:    "Comparing " + strconv.Itoa(arr[j]) + " and " + strconv.Itoa(arr[j+1]),
			})
			stepNum++

			if arr[j] > arr[j+1] {
				// Swap
				arr[j], arr[j+1] = arr[j+1], arr[j]
				result.Swaps++
				swapped = true

				result.Steps = append(result.Steps, Step{
					Array:      append([]int(nil), arr...),
					Comparing:  []int{},
					Swapped:    []int{j, j + 1},
					Sorted:     getSortedIndices(sorted),
					StepNumber: stepNum,
					Message:    "Swapped positions " + strconv.Itoa(j) + " and " + strconv.Itoa(j+1),
				})
				stepNum++
			}
		}

		sorted[n-i-1] = true

		if !swapped {
			result.Steps = append(result.Steps, Step{
				Array:      append([]int(nil), arr...),
				Comparing:  []int{},
				Swapped:    []int{},
				Sorted:     getSortedIndices(sorted),
				StepNumber: stepNum,
				Message:    "Array is sorted!",
			})
			break
		}
	}

	// Mark all as sorted
	for i := 0; i < n; i++ {
		sorted[i] = true
	}

	result.Steps = append(result.Steps, Step{
		Array:      append([]int(nil), arr...),
		Comparing:  []int{},
		Swapped:    []int{},
		Sorted:     getSortedIndices(sorted),
		StepNumber: stepNum,
		Message:    "Sorting complete!",
	})

	return result
}

func getSortedIndices(sorted []bool) []int {
	result := make([]int, 0)
	for i, v := range sorted {
		if v {
			result = append(result, i)
		}
	}
	return result
}

// API Handlers

func handleSort(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	if r.Method == "OPTIONS" {
		w.Header().Set("Access-Control-Allow-Methods", "POST, OPTIONS")
		w.Header().Set("Access-Control-Allow-Headers", "Content-Type")
		return
	}

	var req struct {
		Array     []int  `json:"array"`
		Algorithm string `json:"algorithm"`
	}

	if err := json.NewDecoder(r.Body).Decode(&req); err != nil {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Invalid request"})
		return
	}

	if len(req.Array) == 0 {
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(map[string]string{"error": "Array cannot be empty"})
		return
	}

	var result SortResult

	switch strings.ToLower(req.Algorithm) {
	case "bubble":
		result = BubbleSort(req.Array)
	default:
		result = BubbleSort(req.Array)
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(result)
}

func handleGenerateArray(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	w.Header().Set("Access-Control-Allow-Origin", "*")

	size := 20
	if s := r.URL.Query().Get("size"); s != "" {
		if parsed, err := strconv.Atoi(s); err == nil && parsed > 0 && parsed <= 100 {
			size = parsed
		}
	}

	arr := make([]int, size)
	for i := 0; i < size; i++ {
		arr[i] = (i * 7) % 100 // Pseudo-random but deterministic
	}

	// Fisher-Yates shuffle
	for i := len(arr) - 1; i > 0; i-- {
		j := (i * 13) % (i + 1)
		arr[i], arr[j] = arr[j], arr[i]
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(map[string]interface{}{"array": arr})
}

func serveStatic(dir string) http.Handler {
	return http.FileServer(http.Dir(dir))
}

func main() {
	router := mux.NewRouter()

	// API endpoints
	router.HandleFunc("/api/sort", handleSort).Methods("POST", "OPTIONS")
	router.HandleFunc("/api/generate", handleGenerateArray).Methods("GET", "OPTIONS")

	// Serve static files (prefer Next.js export at frontend/out)
	nextOut := filepath.Join(".", "frontend", "out")
	viteDist := filepath.Join(".", "frontend", "dist")
	if _, err := os.Stat(nextOut); err == nil {
		router.PathPrefix("/").Handler(serveStatic(nextOut))
	} else if _, err := os.Stat(viteDist); err == nil {
		router.PathPrefix("/").Handler(serveStatic(viteDist))
	} else {
		// If dist doesn't exist, serve a simple fallback
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
	<p>Build the frontend with: npm run build</p>
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

