package main

import "fmt"

/*
Ejercicio 2 - Productos
Algunas tiendas ecommerce necesitan realizar una funcionalidad en Go para administrar productos y retornar el valor del precio total.
La empresa tiene 3 tipos de productos: Small, Medium y Large. (Se espera que sean muchos más)
Y los costos adicionales son:
Small: solo tiene el costo del producto
Medium: el precio del producto + un 3% de mantenerlo en la tienda
Large: el precio del producto + un 6% de mantenerlo en la tienda, y adicional a eso $2500 de costo de envío.
El porcentaje de mantenerlo en la tienda es en base al precio del producto.
El costo de mantener el producto en stock en la tienda es un porcentaje del precio del producto.
Se requiere una función factory que reciba el tipo de producto y el precio y retorne una interfaz Product que tenga el método Price.
Se debe poder ejecutar el método Price y que el método me devuelva el precio total en base al costo del producto y los adicionales
en caso que los tenga
*/

type Product interface {
	Price() float64
}

type SmallProduct struct {
	price float64
}

func (s SmallProduct) Price() float64 {
	return s.price
}

type MediumProduct struct {
	price float64
}

func (m MediumProduct) Price() float64 {
	return m.price * 1.03
}

type LargeProduct struct {
	price float64
}

func (l LargeProduct) Price() float64 {
	return l.price*1.06 + 2500
}

func ProductFactory(productType string, price float64) Product {
	switch productType {
	case "Small":
		return SmallProduct{price: price}
	case "Medium":
		return MediumProduct{price: price}
	case "Large":
		return LargeProduct{price: price}
	default:
		panic("Invalid product type")
	}
}

func main() {
	productType := "Large"
	price := 5000.0
	product := ProductFactory(productType, price)

	fmt.Printf("The price of a %s product is $%.2f\n", productType, product.Price())

}
