package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	// ================= soal 1 =================
	fmt.Println("\n================= JAWABAN 1 =================")

	fmt.Println("=== Segitiga Sama Sisi ===")
	var segitiga hitungBangunDatar = segitigaSamaSisi{20, 10}
	fmt.Printf("Luas     : %d cm2\n", segitiga.luas())
	fmt.Printf("Keliling : %d cm\n\n", segitiga.keliling())

	fmt.Println("=== Persegi Panjang ===")
	var persegiPanjang hitungBangunDatar = persegiPanjang{18, 12}
	fmt.Printf("Luas     : %d cm2\n", persegiPanjang.luas())
	fmt.Printf("Keliling : %d cm\n\n", persegiPanjang.keliling())

	fmt.Println("=== Tabung ===")
	var tabung hitungBangunRuang = tabung{7.0, 28.0}
	fmt.Printf("Volume         : %s cm3\n", strings.TrimRight(strings.TrimRight(fmt.Sprintf("%.2f", tabung.volume()), "0"), "."))
	fmt.Printf("Luas Permukaan : %s cm2\n\n", strings.TrimRight(strings.TrimRight(fmt.Sprintf("%.2f", tabung.luasPermukaan()), "0"), "."))

	fmt.Println("=== Balok ===")
	var balok hitungBangunRuang = balok{4, 3, 2}
	fmt.Printf("Volume         : %s cm3\n", strings.TrimRight(strings.TrimRight(fmt.Sprintf("%.2f", balok.volume()), "0"), "."))
	fmt.Printf("Luas Permukaan : %s cm2\n", strings.TrimRight(strings.TrimRight(fmt.Sprintf("%.2f", balok.luasPermukaan()), "0"), "."))

	// ================= soal 2 =================
	fmt.Println("\n================= JAWABAN 2 =================")
	var phone phoneInterface = phone{
		name:   "Samsung Galaxy Note 20",
		brand:  "Samsung Galaxy Note 20",
		year:   2020,
		colors: []string{"Mystic Bronze", "Mystic White", "Mystic Black"},
	}
	fmt.Println(phone.cetakInfoPhone())

	// ================= soal 3 =================
	fmt.Println("\n================= JAWABAN 3 =================")
	fmt.Println(luasPersegi(4, true))
	fmt.Println(luasPersegi(8, false))
	fmt.Println(luasPersegi(0, true))
	fmt.Println(luasPersegi(0, false))

	// ================= soal 4 =================
	fmt.Println("\n================= JAWABAN 4 =================")
	var prefix interface{} = "hasil penjumlahan dari "
	var kumpulanAngkaPertama interface{} = []int{6, 8}
	var kumpulanAngkaKedua interface{} = []int{12, 14}

	s := prefix.(string)
	total := 0
	processOutput(&s, &total, kumpulanAngkaPertama, kumpulanAngkaKedua)
	processOutput(&s, &total, kumpulanAngkaKedua, nil)

	fmt.Println(s)
}

// ================================================
// Struct soal 1
type segitigaSamaSisi struct {
	alas, tinggi int
}

type persegiPanjang struct {
	panjang, lebar int
}

type tabung struct {
	jariJari, tinggi float64
}

type balok struct {
	panjang, lebar, tinggi int
}

// Interface soal 1
type hitungBangunDatar interface {
	luas() int
	keliling() int
}

type hitungBangunRuang interface {
	volume() float64
	luasPermukaan() float64
}

// Function soal 1
func (s segitigaSamaSisi) luas() int {
	return (1 * s.alas * s.tinggi) / 2
}
func (s segitigaSamaSisi) keliling() int {
	return s.alas * 3
}

func (pp persegiPanjang) luas() int {
	return pp.panjang * pp.lebar
}
func (pp persegiPanjang) keliling() int {
	return (2 * pp.panjang) + (2 * pp.lebar)
}

func (t tabung) volume() float64 {
	return float64((22 * math.Pow(t.jariJari, 2) * t.tinggi) / 7)
}
func (t tabung) luasPermukaan() float64 {
	return (2 * 22 * float64(t.jariJari) * (float64(t.jariJari + t.tinggi))) / 7
}

func (b balok) volume() float64 {
	return float64(b.panjang * b.lebar * b.tinggi)
}
func (b balok) luasPermukaan() float64 {
	return float64(2 * ((b.panjang * b.lebar) + (b.lebar * b.tinggi) + (b.panjang * b.tinggi)))
}

// ================================================
// Struct soal 2
type phone struct {
	name, brand string
	year        int
	colors      []string
}

// Interface soal 2
type phoneInterface interface {
	cetakInfoPhone() string
}

// Method soal 2
func (p phone) cetakInfoPhone() string {
	s := "name : " + p.name + "\n" + "brand : " + p.brand + "\n" + "year : " + strconv.FormatInt(int64(p.year), 10) + "\n" + "colors : "
	for i, color := range p.colors {
		s += color

		if i != len(p.colors)-1 {
			s += ", "
		}
	}

	return s
}

// ================================================
// Function soal 3
func luasPersegi(value int, status bool) (text interface{}) {
	if status {
		if value == 0 {
			text = "Maaf anda belum menginput sisi dari persegi"
		} else {
			text = int32(math.Pow(float64(value), 2))
		}
	} else {
		if value == 0 {
			text = nil
		} else {
			text = value
		}
	}

	return
}

// ================================================
// Function soal 4
func processOutput(s *string, total *int, arrPertama interface{}, arrKedua interface{}) {
	for i, angka := range arrPertama.([]int) {
		*total += angka
		*s += strconv.FormatInt(int64(angka), 10)

		if i != len(arrPertama.([]int))-1 {
			*s += " + "
		} else {
			if arrKedua != nil && len(arrKedua.([]int)) != 0 {

				*s += " + "
			} else {
				*s += " = " + strconv.FormatInt(int64(*total), 10)
			}
		}
	}
}
