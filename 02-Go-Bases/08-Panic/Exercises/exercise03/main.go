package main

import (
	"errors"
	"fmt"
)

/*
Ejercicio 3 - Registrando clientes
El mismo estudio del ejercicio anterior, solicita una funcionalidad para poder registrar datos de nuevos clientes.
Los datos requeridos son:

File
Name
ID
Phone number
Home

Tarea 1: Antes de registrar a un cliente, debés verificar si el mismo ya existe. Para ello, necesitás leer los
datos de un array. En caso de que esté repetido, debes manipular adecuadamente el error como hemos visto hasta aquí.
Ese error deberá:
1.- generar un panic;
2.- lanzar por consola el mensaje: “Error: client already exists”, y continuar con la ejecución del programa
normalmente.
Tarea 2: Luego de intentar verificar si el cliente a registrar ya existe, desarrollá una función para validar
que todos los datos a registrar de un cliente contienen un valor distinto de cero. Esta función debe retornar,
al menos, dos valores. Uno de ellos tendrá que ser del tipo error para el caso de que se ingrese por parámetro
algún valor cero (recordá los valores cero de cada tipo de dato, ej: 0, “”, nil).
Tarea 3: Antes de finalizar la ejecución, incluso si surgen panics, se deberán imprimir por consola los
siguientes mensajes: “End of execution” y “Several errors were detected at runtime”. Utilizá defer para cumplir
con este requerimiento.
Requerimientos generales:
Utilizá recover para recuperar el valor de los panics que puedan surgir
Recordá realizar las validaciones necesarias para cada retorno que pueda contener un valor error.
Generá algún error, personalizandolo a tu gusto utilizando alguna de las funciones de Go (realiza también la
	validación pertinente para el caso de error retornado).
*/

type Customer struct {
	id          int
	name        string
	phoneNumber string
	home        string
	file        string
}

var customers = []Customer{}

func validateCustomer(customer Customer) (bool, error) {
	if customer.id == 0 {
		return false, errors.New("The id is required")
	}
	if customer.name == "" {
		return false, errors.New("The name is required")
	}
	if customer.phoneNumber == "" {
		return false, errors.New("The phone number is required")
	}
	if customer.home == "" {
		return false, errors.New("The home is required")
	}
	if customer.file == "" {
		return false, errors.New("The file is required")
	}
	return true, nil
}

func addCustomer(customer Customer) {
	for _, c := range customers {
		if c.id == customer.id {
			panic("Error: client already exists")
		}
	}

	ok, _ := validateCustomer(customer)
	if !ok {
		panic("Error: some data is missing")
	}

	customers = append(customers, customer)
}

func main() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Several errors were detected at runtime: %v\n", r)
		}
		fmt.Println("End of execution")
	}()

	customer1 := Customer{
		id:          1,
		name:        "John Doe",
		phoneNumber: "123456789",
		home:        "Street 123",
		file:        "file1.txt",
	}

	customer2 := Customer{
		id:          2,
		name:        "Jane Doe",
		phoneNumber: "987654321",
		home:        "Street 321",
		file:        "file2.txt",
	}

	//customer3 := customer1

	addCustomer(customer1)
	addCustomer(customer2)
	//addCustomer(customer3)

}
