package exercise01

func CalculateTax(salary float64) float64 {
	var tax float64
	if salary > 150000 {
		tax = salary * 0.27
	} else if salary > 50000 {
		tax = salary * 0.17
	}
	return tax
}
