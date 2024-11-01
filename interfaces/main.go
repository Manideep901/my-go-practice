package main

import "log"

type MototVehicle interface {
	Mileage() float64
}

type BMW struct {
	distance     float64
	fuel         float64
	averageSpeed string
}

type Audi struct {
	distance     float64
	fuel         float64
	averageSpeed string
}

func main() {
	b1 := BMW{
		distance:     167.9,
		fuel:         36,
		averageSpeed: "58",
	}
	a1 := Audi{
		distance:     78.9,
		fuel:         30,
		averageSpeed: "45",
	}
	person := []MototVehicle{b1, a1}
	totalMileage(person)
}

func (b BMW) Mileage() float64 {
	return b.distance / b.fuel
}

func (a Audi) Mileage() float64 {
	return a.distance / a.fuel
}

func totalMileage(m []MototVehicle) {
	tm := 0.0

	for _, v := range m {
		tm = tm + v.Mileage()
	}
	log.Printf("Total mileage per month %f km/ltr", tm)
}
