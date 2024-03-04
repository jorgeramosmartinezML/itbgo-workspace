package main

import (
	"errors"
	"fmt"
)

/*
Ejercicio 3 - Impuestos de salario #3
Hacé lo mismo que en el ejercicio anterior pero reformulando el código para que,
en reemplazo de “Error()”,  se implemente “errors.New()”.
*/

var SalaryError = errors.New("Salary is less than 10000")

func main() {

	salary := 1000
	if salary <= 10000 {
		err := fmt.Errorf("New Error: %w", SalaryError)
		if errors.Is(err, SalaryError) {
			fmt.Println(err.Error())
		}
	}

}
