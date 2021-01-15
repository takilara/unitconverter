package main

import (
	"math"
)

func ToSi(val float64, uom string) float64 {
	
	if math.IsNaN(val) || uom == "" {
		return math.NaN()
	} else {
		return 0.0
	}
}

func FromSi(val float64, uom string) float64 {
        if math.IsNaN(val) || uom == "" {
                return math.NaN()
        } else {
                return 0.0
        }
}
