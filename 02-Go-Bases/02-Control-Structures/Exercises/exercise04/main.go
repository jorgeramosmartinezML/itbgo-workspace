package main

import "fmt"

/*
=> Ejercicio 4
Un empleado de una empresa quiere saber el nombre y edad de uno de sus empleados. Según el siguiente mapa,
ayudá a imprimir la edad de Benjamin.

var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

Por otro lado también es necesario:
Saber cuántos de sus empleados son mayores de 21 años.
Agregar un empleado nuevo a la lista, llamado Federico que tiene 25 años.
Eliminar a Pedro del mapa.
*/

func main() {

	var employees = map[string]int{"Benjamin": 20, "Nahuel": 26, "Brenda": 19, "Darío": 44, "Pedro": 30}

	// 1.- To know the name and age of the employees
	for name, age := range employees {
		fmt.Println(name, age)
	}

	// 2.- To know how many employees are over 21 years old
	var over21 int
	for _, age := range employees {
		if age > 21 {
			over21++
		}
	}
	fmt.Println("There are", over21, "employees over 21 years old")

	// 3.- To add a new employee to the list
	employees["Federico"] = 25
	fmt.Println(employees)

	// 4.- To delete Pedro from the map
	delete(employees, "Pedro")
	fmt.Println(employees)

}
