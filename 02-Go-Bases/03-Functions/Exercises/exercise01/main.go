package main

import "fmt"

/*
Una empresa de chocolates necesita calcular el impuesto de sus empleados al momento de depositar el sueldo,
para cumplir el objetivo es necesario crear una función que devuelva el impuesto de un salario.

Teniendo en cuenta que si la persona gana más de $50.000 se le descontará un 17 % del sueldo y
si gana más de $150.000 se le descontará además un 10 % (27% en total).
*/

func CalculateTax(salary float64) float64 {
	var tax float64
	if salary > 150000 {
		tax = salary * 0.27
	} else if salary > 50000 {
		tax = salary * 0.17
	}
	return tax
}

func main() {
	salary1 := 160000.0
	tax1 := CalculateTax(salary1)
	fmt.Println("The tax for", salary1, "is", tax1)

	salary2 := 60000.0
	tax2 := CalculateTax(salary2)
	fmt.Println("The tax for", salary2, "is", tax2)

	salary3 := 30000.0
	tax3 := CalculateTax(salary3)
	fmt.Println("The tax for", salary3, "is", tax3)
}
