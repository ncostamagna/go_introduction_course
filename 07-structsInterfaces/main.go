package main

import (
	"encoding/json"
	"fmt"

	"github.com/ncostamagna/go_introduction_course/07-structsInterfaces/structsInterface/structs"
	"github.com/ncostamagna/go_introduction_course/07-structsInterfaces/structsInterface/vehicles"
)

func main() {

	var p1 structs.Product
	p1.ID = 12
	p1.Name = "test"
	fmt.Println(p1)
	fmt.Println()

	p2 := structs.Product{}
	p2.ID = 12
	p2.Name = "test"
	fmt.Println(p2)
	fmt.Println()

	p3 := structs.Product{2, "test 3", structs.Type{}, 0, 12.21, nil}
	fmt.Println(p3)
	fmt.Println()

	p4 := structs.Product{
		ID:   3,
		Name: "My product",
	}
	fmt.Println(p4)
	fmt.Println()

	p5 := structs.Product{
		Name: "Heladera marca sarasa",
		Type: structs.Type{
			Code:        "A",
			Description: "Electrodomestico",
		},
		Tags: []string{"heladera", "sarasa", "freezer", "refrigerador"},
	}
	fmt.Println(p5)
	fmt.Println()

	p6 := structs.Product{
		Name:  "Heladera marca sarasa",
		Price: 40000,
		Type: structs.Type{
			Code:        "A",
			Description: "Electrodomestico",
		},
		Tags:  []string{"heladera", "sarasa", "freezer", "refrigerador"},
		Count: 5,
	}

	v, err := json.Marshal(p6)
	fmt.Println(err)
	fmt.Println(string(v))

	fmt.Println("Precio total: ", p6.TotalPrice())
	fmt.Println()
	fmt.Println(p6)
	p6.SetName("other name")
	p6.AddTags("tag1", "tag2", "tag3")
	fmt.Println(p6)
	fmt.Println()

	p7 := structs.Product{
		Name:  "Naranja",
		Price: 20,
		Type: structs.Type{
			Code:        "B",
			Description: "Alimento",
		},
		Tags:  []string{"alimento", "verdura"},
		Count: 20,
	}

	p8 := structs.Product{
		Name:  "Cortina",
		Price: 2700,
		Type: structs.Type{
			Code:        "C",
			Description: "Hogar",
		},
		Tags:  []string{"hogar", "cortina"},
		Count: 3,
	}

	car := structs.NewCar(11212)
	car.AddProducts(p6, p7, p8)

	fmt.Println("PRODUCTS CAR")
	fmt.Println("Total Products: ", len(car.Products))
	fmt.Printf("Total Car: $%.2f\n", car.Total())

	fmt.Println()
	fmt.Println("VEHICLES")

	carV := vehicles.Car{Time: 160}
	fmt.Println(carV.Distance())

	vArray := []string{"CAR", "MOTORCYCLE", "TRUCK", "MOTORCYCLE", "TRUCK", "GOKU", "SARASA"}

	var d float64
	for _, v := range vArray {
		fmt.Printf("Vechile: %s\n", v)
		vech, err := vehicles.New(v, 400)
		if err != nil {
			fmt.Println("Error: ", err)
			fmt.Println()
			continue
		}
		distance := vech.Distance()
		fmt.Printf("Distance: %.2f\n", distance)
		fmt.Println()
		d += distance
	}
	fmt.Printf("Total distance: %.2f\n", d)
	fmt.Println()
}
