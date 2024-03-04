package exercise03

var (
	CategoryC float64 = 1000
	CategoryB float64 = 1500
	CategoryA float64 = 3000
)

func CalculateSalary(minutesWorked int, category string) float64 {
	var salary float64
	switch category {
	case "C":
		salary = 1000 * (float64(minutesWorked) / 60)
	case "B":
		salary = 1500 * (float64(minutesWorked) / 60)
		extra := salary * 0.2
		salary += extra
	case "A":
		salary = 3000 * (float64(minutesWorked) / 60)
		extra := salary * 0.5
		salary += extra
	}
	return salary
}
