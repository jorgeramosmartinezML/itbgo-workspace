package exercise01

/*
Ejercicio 1 - Testear el impuesto del salario
La  empresa de chocolates que anteriormente necesitaba calcular el
impuesto de sus empleados al momento de depositar el sueldo de los
mismos ahora nos solicitó validar que los cálculos de estos impuestos
están correctos. Para esto nos encargaron el trabajo de realizar los
test correspondientes para:

Calcular el impuesto en caso de que el empleado gane por debajo de $50.000.
Calcular el impuesto en caso de que el empleado gane por encima de $50.000.
Calcular el impuesto en caso de que el empleado gane por encima de $150.000.
*/

import "testing"

func TestCalculateTax_less50000(t *testing.T) {
	// Arrange
	salary := 40000.0
	expectedTax := 0.0

	// Act
	tax := CalculateTax(salary)

	// Assert
	if tax != expectedTax {
		t.Errorf("Expected tax %f, got %f", expectedTax, tax)
		return
	}
	t.Log("Test passed")
}

func TestCalculateTax_over50000(t *testing.T) {
	// Arrange
	salary := 60000.0
	expectedTax := 10200.0

	// Act
	tax := CalculateTax(salary)

	// Assert
	if tax != expectedTax {
		t.Errorf("Expected tax %f, got %f", expectedTax, tax)
		return
	}
	t.Log("Test passed")
}

func TestCalculatorTax_over150000(t *testing.T) {
	// Arrange
	salary := 200000.0
	expectedTax := 54000.0

	// Act
	tax := CalculateTax(salary)

	// Assert
	if tax != expectedTax {
		t.Errorf("Expected tax %f, got %f", expectedTax, tax)
		return
	}
	t.Log("Test passed")
}
