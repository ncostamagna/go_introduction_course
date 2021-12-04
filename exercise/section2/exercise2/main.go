package main

import "github.com/ncostamagna/go_introduction_course/exercise/section2/exercise2/commerce"

/*
Un ecommerce llamado Roquetete necesita implementar una funcionalidad para administrar sus productos.
La empresa tiene 3 tipos de productos:
 - Pequeño: El costo del producto (sin costo adicional)
 - Mediano: El costo del producto + un 3% por mantenerlo en el almacén de la tienda.
 - Grande: El costo del producto + un 6%  por mantenimiento, y un costo adicional  por envío de $2500.

Necesita generar una estructura commerce que reciba los productos y tenga una funcionalidad para imprimir el subtotal de cada producto y el total
*/

func main() {
	p1 := commerce.NewProduct(commerce.Small, "pinza", 5, 20)
	p2 := commerce.NewProduct(commerce.Small, "clavo", 20, 5)
	p3 := commerce.NewProduct(commerce.Medium, "microondas", 1, 20000)
	p4 := commerce.NewProduct(commerce.Medium, "monitor", 2, 15000)
	p5 := commerce.NewProduct(commerce.Large, "heladera", 1, 60000)

	c := commerce.New()
	c.Add(p1, p2, p3, p4, p5)
	c.Print()
}
