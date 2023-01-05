package models

// Struct
type Balok struct {
	Panjang, Lebar, Tinggi int
}

// Method
func (b Balok) Volume() float64 {
	return float64(b.Panjang * b.Lebar * b.Tinggi)
}

func (b Balok) LuasPermukaan() float64 {
	return float64(2 * ((b.Panjang * b.Lebar) + (b.Lebar * b.Tinggi) + (b.Panjang * b.Tinggi)))
}
