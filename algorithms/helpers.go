package algorithms

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

func getSortedIndices(sorted []bool) []int {
	result := make([]int, 0)
	for i, v := range sorted {
		if v {
			result = append(result, i)
		}
	}
	return result
}

func HandleSort(w http.ResponseWriter, r *http.Request) {
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

func HandleGenerateArray(w http.ResponseWriter, r *http.Request) {
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
