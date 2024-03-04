package exercise01

import "fmt"

/*
Ejercicio 1 - Registro de estudiantes
Una universidad necesita registrar a los/as estudiantes y generar una funcionalidad para imprimir el
detalle de los datos de cada uno de ellos/as, de la siguiente manera:


Name: [Nombre del alumno]
Apellido: [Apellido del alumno]
DNI: [DNI del alumno]
Fecha: [Fecha ingreso alumno]

Los valores que están en corchetes deben ser reemplazados por los datos brindados por los alumnos/as.
Para ello es necesario generar una estructura Alumnos con las variables Nombre, Apellido, DNI, Fecha y que tenga un método detalle
*/

type Student struct {
	Name     string
	LastName string
	DNI      string
}

func (s Student) Detail() string {
	return fmt.Sprintf("Name: %s\nLastName: %s\nDNI: %s\n", s.Name, s.LastName, s.DNI)
}

func main() {
	name := "Juan"
	lastName := "Perez"
	DNI := "12345678"

	student1 := Student{name, lastName, DNI}
	student1.Detail()
}
