package functions

import (
	"math"
)

func Volume(JariJari int, Tinggi int) float64 {
	return math.Pi * math.Pow(float64(JariJari), 2) * float64(Tinggi)
}

func LuasAlas(JariJari int) float64 {
	return math.Pi * math.Pow(float64(JariJari), 2)
}

func KelilingAlas(JariJari int) float64 {
	return 2 * math.Pi * float64(JariJari)
}
