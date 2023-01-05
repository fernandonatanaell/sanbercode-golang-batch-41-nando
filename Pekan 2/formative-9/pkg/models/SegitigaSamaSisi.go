package models

// Struct
type SegitigaSamaSisi struct {
	Alas, Tinggi int
}

// Method
func (s SegitigaSamaSisi) Luas() int {
	return (1 * s.Alas * s.Tinggi) / 2
}

func (s SegitigaSamaSisi) Keliling() int {
	return s.Alas * 3
}
