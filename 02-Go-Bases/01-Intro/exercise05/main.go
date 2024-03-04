package main

func main() {

	// Incorrecto: no se puede asignar un valor de tipo int a una variable de tipo string
	// var student1 string = 23
	// Correcto
	var student1 int = 23

	// Incorrecto: no se puede asignar un valor de tipo string a una variable de tipo float64
	// var grade float64 = "A"
	// Correcto
	var grade float64 = 10.5

	// Correcto
	var isEnrolled = "yes"

	// Incorrecto: no se puede utilizar el operador de asignación := para declarar una variable que ya existe
	// var number_of_students int
	// number_of_students := 5
	// Correcto
	var number_of_students int
	number_of_students = 5

	println(student1, grade, isEnrolled, number_of_students)
}
