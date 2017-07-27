package vcl

type IVehicle interface {
	Range() float64           // the range in km that the vehicle can travel
	Drive(km float64) float64 // drive the given km and return the travelled distance
}

// Vehicle moves
type Vehicle struct {
	Fuel   float64 // fuel in liters
	AvgKph float64
	AvgKpl float64
}

func (v *Vehicle) Range() float64 {
	return v.Fuel * v.AvgKpl
}

func (v *Vehicle) Drive(km float64) float64 {
	r := v.Range()
	if km < r {
		v.Fuel = v.Fuel - km/v.AvgKpl
		return km
	}

	v.Fuel = 0
	return r
}

type Car struct {
	Vehicle
}

type Plane struct {
	Vehicle
}
