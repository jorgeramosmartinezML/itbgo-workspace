package main

/*
Ejercicio 1
Crear un programa que cumpla los siguiente puntos:

Tener una estructura llamada Product con los campos ID, Name, Price, Description y Category.
Tener un slice global de Product llamado Products instanciado con valores.
2 métodos asociados a la estructura Product: Save(), GetAll(). El método Save() deberá tomar el slice de Products y añadir el producto desde el cual se llama al método. El método GetAll() deberá imprimir todos los productos guardados en el slice Products.
Una función getById() al cual se le deberá pasar un INT como parámetro y retorna el producto correspondiente al parámetro pasado.
Ejecutar al menos una vez cada método y función definido desde main().
*/

import "fmt"

type Product struct {
	Id          int
	Name        string
	Price       float32
	Description string
	Category    string
}

var product1 = Product{1, "Laptop", 1000, "Laptop Dell", "Electronics"}

var Products = []Product{product1}

func (p Product) Save() {
	Products = append(Products, p)
}

func (p Product) GetAll() {
	for _, product := range Products {
		fmt.Println(product)
	}
}

func getById(id int) Product {
	for _, product := range Products {
		if product.Id == id {
			return product
		}
	}
	return Product{}
}

func main() {

	productA := Product{2, "Celular", 500, "Celular Samsung", "Electronics"}
	productA.Save()
	productA.GetAll()

	productB := Product{3, "Playera", 100, "Playera Nike", "Ropa"}
	productB.Save()
	productB.GetAll()

	productC := getById(1)
	fmt.Println(productC)

}
