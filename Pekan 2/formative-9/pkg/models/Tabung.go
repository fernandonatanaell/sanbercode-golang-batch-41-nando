package models

import "math"

// Struct
type Tabung struct {
	JariJari, Tinggi float64
}

// Method
func (t Tabung) Volume() float64 {
	return float64((22 * math.Pow(t.JariJari, 2) * t.Tinggi) / 7)
}

func (t Tabung) LuasPermukaan() float64 {
	return (2 * 22 * float64(t.JariJari) * (float64(t.JariJari + t.Tinggi))) / 7
}
