package algorithms

type Step struct {
	Array      []int  `json:"array"`
	Comparing  []int  `json:"comparing"`
	Swapped    []int  `json:"swapped"`
	Sorted     []int  `json:"sorted"`
	StepNumber int    `json:"step_number"`
	Message    string `json:"message"`
}

// SortResult contains all steps for a complete sort
type SortResult struct {
	Steps       []Step `json:"steps"`
	Algorithm   string `json:"algorithm"`
	Duration    string `json:"duration"`
	Comparisons int    `json:"comparisons"`
	Swaps       int    `json:"swaps"`
}
