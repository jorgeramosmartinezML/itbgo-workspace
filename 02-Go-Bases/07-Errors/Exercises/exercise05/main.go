package main

import (
	"errors"
	"fmt"
)

/*
Ejercicio 5 -  Impuestos de salario #5
Vamos a hacer que nuestro programa sea un poco más complejo y útil.
Desarrollá las funciones necesarias para permitir a la empresa calcular:
Salario mensual de un trabajador según la cantidad de horas trabajadas.
La función recibirá las horas trabajadas en el mes y el valor de la hora como argumento.
Dicha función deberá retornar más de un valor (salario calculado y error).
En caso de que el salario mensual sea igual o superior a $150.000, se le deberá
descontar el 10 % en concepto de impuesto.
En caso de que la cantidad de horas mensuales ingresadas sea menor a 80 o un número negativo,
la función debe devolver un error. El mismo tendrá que indicar “Error: the worker cannot have
worked less than 80 hours per month”.
*/

func calculateSalary(hours int, valueHour int) (int, error) {
	if hours < 80 {
		return 0, errors.New("Error: the worker cannot have worked less than 80 hours per month")
	}
	salary := hours * valueHour
	if salary >= 150000 {
		salary = salary - (salary * 10 / 100)
	}
	return salary, nil
}

func main() {

	hours := 20
	valueHour := 1000
	salary, err := calculateSalary(hours, valueHour)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Printf("The salary of the worker is: %d\n", salary)

}
