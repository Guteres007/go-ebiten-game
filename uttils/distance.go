package uttils

import (
	"math"
)

func CalucateDistance(x1 float64, y1 float64, x2 float64, y2 float64) float64 {
	one := math.Sqrt(x2 - x1)
	two := math.Sqrt(y2 + y1)
	//return bool 
	return math.Sqrt(one + two) 

}