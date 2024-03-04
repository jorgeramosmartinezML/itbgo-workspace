package main

import "fmt"

/*
Una empresa marinera necesita calcular el salario de sus empleados basándose en la cantidad de horas trabajadas por mes y la categoría.

Categoría C, su salario es de $1.000 por hora.
Categoría B, su salario es de $1.500 por hora, más un 20 % de su salario mensual.
Categoría A, su salario es de $3.000 por hora, más un 50 % de su salario mensual.
Se solicita generar una función que reciba por parámetro la cantidad de minutos trabajados por mes, la categoría y que devuelva su salario.
*/

var (
	CategoryC float64 = 1000
	CategoryB float64 = 1500
	CategoryA float64 = 3000
)

func calculateSalary(minutesWorked int, category string) float64 {
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

func main() {

	minutesWorked := 180
	category := "C"
	salary := calculateSalary(minutesWorked, category)
	fmt.Println("The salary is: $", salary, "for", minutesWorked, "minutes worked and category", category)

	minutesWorked = 500
	category = "B"
	salary = calculateSalary(minutesWorked, category)
	fmt.Println("The salary is: $", salary, "for", minutesWorked, "minutes worked and category", category)

	minutesWorked = 1000
	category = "A"
	salary = calculateSalary(minutesWorked, category)
	fmt.Println("The salary is: $", salary, "for", minutesWorked, "minutes worked and category", category)

}
