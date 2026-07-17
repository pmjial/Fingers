package convert

import "Fingers/db"

// Unit values represent their conversion to meters
type Unit = float64

const (
	Inch  Unit = 0.0254
	Foot  Unit = 0.3048
	Yard  Unit = 0.9144
	Mile  Unit = 1610
	Cm    Unit = 0.01
	Meter Unit = 1
	Km    Unit = 1000
)

func InMeter(length db.Length) db.Length {
	return 1 / length
}
