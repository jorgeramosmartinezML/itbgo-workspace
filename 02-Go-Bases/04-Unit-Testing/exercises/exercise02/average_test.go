package exercise02

/*
Ejercicio 2 - Calcular promedio
El colegio informó que las operaciones para calcular el promedio no
se están realizando correctamente, por lo que ahora nos corresponde
realizar los test correspondientes:
Calcular el promedio de las notas de los alumnos.
*/

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestAverage(t *testing.T) {
	// Arrange
	notes := []int{1, 2, 3, 4, 5}
	expectedAverage := 3.0

	// Act
	average := Average(notes...)

	// Assert
	require.Equal(t, expectedAverage, average)
}
