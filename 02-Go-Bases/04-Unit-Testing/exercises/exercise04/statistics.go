package exercise04

import "fmt"

const (
	minimum = "minimum"
	average = "average"
	maximum = "maximum"
)

func Min(numbers ...int) int {
	minValue := numbers[0]
	for _, number := range numbers {
		if number < minValue {
			minValue = number
		}
	}
	return minValue
}

func Max(numbers ...int) int {
	maxValue := numbers[0]
	for _, number := range numbers {
		if number > maxValue {
			maxValue = number
		}
	}
	return maxValue
}

func Average(numbers ...int) int {
	sum := 0
	for _, number := range numbers {
		sum += number
	}
	return sum / len(numbers)
}

func Operation(operation string) (func(...int) int, error) {
	switch operation {
	case minimum:
		return Min, nil
	case average:
		return Average, nil
	case maximum:
		return Max, nil
	default:
		return nil, fmt.Errorf("operation %s not found", operation)
	}
}
