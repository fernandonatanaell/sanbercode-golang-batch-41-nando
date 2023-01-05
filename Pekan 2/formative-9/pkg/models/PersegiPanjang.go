package models

// Struct
type PersegiPanjang struct {
	Panjang, Lebar int
}

// Method
func (pp PersegiPanjang) Luas() int {
	return pp.Panjang * pp.Lebar
}

func (pp PersegiPanjang) Keliling() int {
	return (2 * pp.Panjang) + (2 * pp.Lebar)
}
