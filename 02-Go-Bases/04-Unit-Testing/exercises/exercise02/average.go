package exercise02

func Average(notes ...int) float64 {
	var sum float64
	for _, note := range notes {

		sum += float64(note)
	}
	average := sum / float64(len(notes))
	return average
}
