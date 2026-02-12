package algorithms

import (
	"bufio"
	"bytes"
	"encoding/json"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"time"
)

type ZigInput struct {
	Array     []int  `json:"array"`
	Algorithm string `json:"algorithm"`
}

func RunZigAlgorithm(algorithm string, arr []int) (SortResult, error) {
	start := time.Now()
	result := SortResult{
		Steps:     make([]Step, 0),
		Algorithm: algorithm,
	}

	// Determine Zig file path
	zigFile := ""
	switch algorithm {
	case "Bubble Sort":
		zigFile = "bubble_sort.zig"
	case "Insertion Sort":
		zigFile = "insertion_sort.zig"
	case "Selection Sort":
		zigFile = "selection_sort.zig"
	default:
		return result, fmt.Errorf("unknown algorithm: %s", algorithm)
	}

	wd, err := os.Getwd()
	if err != nil {
		return result, err
	}

	zigPath := filepath.Join(wd, "tools", "zig", "zig.exe")
	algoPath := filepath.Join(wd, "zig_algorithms", zigFile)

	// Prepare input JSON
	input := ZigInput{
		Array:     arr,
		Algorithm: algorithm,
	}
	inputBytes, err := json.Marshal(input)
	if err != nil {
		return result, err
	}

	// Create command
	cmd := exec.Command(zigPath, "run", algoPath)
	cmd.Stdin = bytes.NewReader(inputBytes)
	
	// Capture stdout/stderr
	var stdout, stderr bytes.Buffer
	cmd.Stdout = &stdout
	cmd.Stderr = &stderr

	// Execute
	if err := cmd.Run(); err != nil {
		return result, fmt.Errorf("zig execution failed: %v, stderr: %s", err, stderr.String())
	}

	// Parse output line by line
	scanner := bufio.NewScanner(&stdout)
	for scanner.Scan() {
		line := scanner.Bytes()
		if len(line) == 0 {
			continue
		}

		var step Step
		if err := json.Unmarshal(line, &step); err != nil {
			// If line isn't JSON, it might be debug output. Log it or ignore.
			// fmt.Printf("Non-JSON output from Zig: %s\n", line)
			continue
		}
		result.Steps = append(result.Steps, step)
		
		// Update stats based on message content or step differences
		// Simplified stats logic:
		if len(step.Swapped) > 0 {
			result.Swaps++
		}
		if len(step.Comparing) > 0 {
			result.Comparisons++
		}
	}

	result.Duration = time.Since(start).String()
	return result, nil
}
