package uttils

import (
	"math"
)

func CalucateDistance(x1 float64, y1 float64, x2 float64, y2 float64) float64 {
    dist :=	math.Sqrt((y2 - y1) * (y2 - y1) + (x2 - x1) * (x2 - x1));
	return dist

}