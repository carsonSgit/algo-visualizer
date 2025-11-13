package algorithms

import (
	"strconv"
)

func BubbleSort(arr []int) SortResult {
	result := SortResult{
		Steps:       make([]Step, 0),
		Algorithm:   "Bubble Sort",
		Swaps:       0,
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
