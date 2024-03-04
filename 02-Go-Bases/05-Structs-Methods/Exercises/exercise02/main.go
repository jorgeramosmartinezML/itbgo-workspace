package main

import "fmt"

/*
Ejercicio 2 - Employee
Una empresa necesita realizar una buena gestión de sus empleados,
para esto realizaremos un pequeño programa nos ayudará a gestionar
correctamente dichos empleados. Los objetivos son:

Crear una estructura Person con los campos ID, Name, DateOfBirth.
Crear una estructura Employee con los campos: ID, Position y una composición con la estructura Person.
Realizar el método a la estructura Employe que se llame PrintEmployee(), lo que hará es realizar la impresión de los campos de un empleado.
Instanciar en la función main() tanto una Person como un Employee cargando sus respectivos campos y por último ejecutar el método PrintEmployee().
Si logras realizar este pequeño programa pudiste ayudar a la empresa a solucionar la gestión de los empleados.
*/

type Person struct {
	id          int
	Name        string
	DateOfBirth string
}

type Employee struct {
	id       int
	Position string
	Person
}

func (e Employee) PrintEmployee() {
	fmt.Printf("The employee %s has the id %d, was born in %s and works as a %s", e.Person.Name, e.id, e.Person.DateOfBirth, e.Position)
}

func main() {
	person1 := Person{1, "Juan", "01/01/2000"}
	employee1 := Employee{1, "Developer", person1}

	employee1.PrintEmployee()
}
