package vehicles

import "fmt"

type Vehicle interface {
	Distance() float64
}

const (
	CarVechicle        = "CAR"
	MotorcycleVechicle = "MOTORCYCLE"
	TruckVehicle       = "TRUCK"
	GokuVehicle        = "GOKU"
)

func New(v string, time int) (Vehicle, error) {
	switch v {
	case CarVechicle:
		return &Car{Time: time}, nil
	case MotorcycleVechicle:
		return &Motorcycle{Time: time}, nil
	case TruckVehicle:
		return &Truck{Time: time}, nil
	case GokuVehicle:
		return &Goku{Time: time}, nil
	}
	return nil, fmt.Errorf("vehicle '%s' not exists", v)
}

type Car struct {
	Time int
}

func (c *Car) Distance() float64 {
	return 100 * (float64(c.Time) / 60)
}

type Motorcycle struct {
	Time int
}

func (c *Motorcycle) Distance() float64 {
	return 120 * (float64(c.Time) / 60)
}

type Truck struct {
	Time int
}

func (c *Truck) Distance() float64 {
	return 70 * (float64(c.Time) / 60)
}

type Goku struct {
	Time int
}

func (c *Goku) Distance() float64 {
	return 300000 * (float64(c.Time) / 60)
}
