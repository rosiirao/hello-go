package main

import (
	"math"
)

type Gis struct {
	longitude float64
	latitude  float64
}

// Earth radius, k
const earth_radius = 6378.137

// -- Query all the points whose distance to 113.914619,22.50128 is less then 2km
// -- radians = d * Math.PI / 180.0; // angles to radians
// -- a = radians(lat1) - radians(lat2)
// -- b = radians(lng1) - radians(lng2)
// -- s = 2 arcsin √(sin²(a/2) + cos(lat1) * cos(lat2) * sin²(b/2)) * 6378.137
// -- 2 * Math.Asin(Math.Sqrt(
// --   Math.Sin(a / 2) * Math.Sin(a / 2) + Math.Cos(radLat1) * Math.Cos(radLat2) * Math.Sin (b / 2) * Math.Sin (b / 2)
// -- )) * 6378.137
func (source *Gis) distance(target *Gis) float64 {
	radians_x, radians_y := source.radians()
	radians_a, radians_b := target.radians()
	a := radians_a - radians_x
	b := radians_b - radians_y
	return math.Asin(
		math.Sqrt(
			math.Pow(math.Sin(a/2), 2)+
				math.Cos(radians_a)*math.Cos(radians_x)*
					math.Pow(math.Sin(b/2), 2))) * 2 * earth_radius
}

func (pos *Gis) radians() (float64, float64) {
	return radians(pos.latitude), radians(pos.longitude)
}

func radians(angle float64) float64 {
	return math.Pi * angle / 180
}
