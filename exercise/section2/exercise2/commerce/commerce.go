package commerce

import "fmt"

type Product interface {
	Total() float64
	Name() string
	Size() string
}

type productSmall struct {
	name  string
	count int
	price float64
}

func (p productSmall) Total() float64 {
	return p.price * float64(p.count)
}

func (p productSmall) Name() string {
	return p.name
}

func (p productSmall) Size() string {
	return string(Small)
}

type productMedium struct {
	name  string
	count int
	price float64
}

func (p productMedium) Total() float64 {
	return (p.price * float64(p.count)) * 1.03
}

func (p productMedium) Name() string {
	return p.name
}

func (p productMedium) Size() string {
	return string(Medium)
}

type productLarge struct {
	name  string
	count int
	price float64
}

func (p productLarge) Total() float64 {
	return (p.price*float64(p.count))*1.06 + 2500
}

func (p productLarge) Name() string {
	return p.name
}

func (p productLarge) Size() string {
	return string(Large)
}

type Size string

const (
	Small  Size = "small"
	Medium Size = "medium"
	Large  Size = "large"
)

func NewProduct(size Size, name string, count int, price float64) Product {

	switch size {
	case Small:
		return &productSmall{count: count, name: name, price: price}
	case Medium:
		return &productMedium{count: count, name: name, price: price}
	case Large:
		return &productLarge{count: count, name: name, price: price}
	}

	return nil
}

type Commerce struct {
	products []Product
}

func New() *Commerce {
	return &Commerce{}
}

func (c *Commerce) Add(p ...Product) {
	c.products = append(c.products, p...)
}

func (c *Commerce) Print() {
	fmt.Println("Products List")
	var total float64
	for _, p := range c.products {
		total += p.Total()
		fmt.Printf("Name: %s, Size: %s, Total: %.2f\n", p.Name(), p.Size(), p.Total())
	}
	fmt.Printf("\nTotal: %.4f\n\n", total)
}
