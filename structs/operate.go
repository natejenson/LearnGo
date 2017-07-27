package main

import (
	"fmt"
	"math/rand"

	"github.com/natejenson/LearnGo/Structs/vcl"
)

func main() {
	c := vcl.Car{Vehicle: vcl.Vehicle{Fuel: 65, AvgKph: 88.5139, AvgKpl: 8.6}}
	p := vcl.Plane{Vehicle: vcl.Vehicle{Fuel: 273, AvgKph: 250, AvgKpl: 5.8}}

	fmt.Printf("Driving the car (Range: %v)...\n", c.Vehicle.Range())
	operate(&c.Vehicle)
	fmt.Printf("Flying the plane (Range: %v)...\n", p.Vehicle.Range())
	operate(&p.Vehicle)
}

func operate(v vcl.IVehicle) {
	i := v.Range() / 4
	t := 0.0
	for v.Range() > 0 {
		r := rand.Float64() * i
		d := v.Drive(r)
		t += d
		fmt.Printf("... travelled %v km, %v total\n", d, t)
	}
	fmt.Println("... out of fuel!")
}
