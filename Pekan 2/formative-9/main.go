package main

import (
	"fmt"
	functions "formative-9/pkg/functions"
	interfaces "formative-9/pkg/interfaces"
	models "formative-9/pkg/models"
	"strings"
)

func main() {
	// ================= SOAL 1 =================
	fmt.Println("\n================= JAWABAN 1 =================")

	fmt.Println("=== Segitiga Sama Sisi ===")
	var segitiga interfaces.HitungBangunDatar = models.SegitigaSamaSisi{Alas: 20, Tinggi: 10}
	fmt.Printf("Luas     : %d cm2\n", segitiga.Luas())
	fmt.Printf("Keliling : %d cm\n\n", segitiga.Keliling())

	fmt.Println("=== Persegi Panjang ===")
	var persegiPanjang interfaces.HitungBangunDatar = models.PersegiPanjang{Panjang: 18, Lebar: 12}
	fmt.Printf("Luas     : %d cm2\n", persegiPanjang.Luas())
	fmt.Printf("Keliling : %d cm\n\n", persegiPanjang.Keliling())

	fmt.Println("=== Tabung ===")
	var tabung interfaces.HitungBangunRuang = models.Tabung{JariJari: 7.0, Tinggi: 28.0}
	fmt.Printf("Volume         : %s cm3\n", strings.TrimRight(strings.TrimRight(fmt.Sprintf("%.2f", tabung.Volume()), "0"), "."))
	fmt.Printf("Luas Permukaan : %s cm2\n\n", strings.TrimRight(strings.TrimRight(fmt.Sprintf("%.2f", tabung.LuasPermukaan()), "0"), "."))

	fmt.Println("=== Balok ===")
	var balok interfaces.HitungBangunRuang = models.Balok{Panjang: 4, Lebar: 3, Tinggi: 2}
	fmt.Printf("Volume         : %s cm3\n", strings.TrimRight(strings.TrimRight(fmt.Sprintf("%.2f", balok.Volume()), "0"), "."))
	fmt.Printf("Luas Permukaan : %s cm2\n", strings.TrimRight(strings.TrimRight(fmt.Sprintf("%.2f", balok.LuasPermukaan()), "0"), "."))

	// ================= SOAL 2 =================
	fmt.Println("\n================= JAWABAN 2 =================")
	var phone interfaces.PhoneInterface = models.Phone{
		Name:   "Samsung Galaxy Note 20",
		Brand:  "Samsung Galaxy Note 20",
		Year:   2020,
		Colors: []string{"Mystic Bronze", "Mystic White", "Mystic Black"},
	}
	fmt.Println(phone.CetakInfoPhone())

	// ================= SOAL 3 =================
	fmt.Println("\n================= JAWABAN 3 =================")
	fmt.Println(functions.LuasPersegi(4, true))
	fmt.Println(functions.LuasPersegi(8, false))
	fmt.Println(functions.LuasPersegi(0, true))
	fmt.Println(functions.LuasPersegi(0, false))

	// ================= SOAL 4 =================
	fmt.Println("\n================= JAWABAN 4 =================")
	var prefix interface{} = "hasil penjumlahan dari "
	var kumpulanAngkaPertama interface{} = []int{6, 8}
	var kumpulanAngkaKedua interface{} = []int{12, 14}

	s := prefix.(string)
	total := 0
	functions.ProcessOutput(&s, &total, kumpulanAngkaPertama, kumpulanAngkaKedua)
	functions.ProcessOutput(&s, &total, kumpulanAngkaKedua, nil)

	fmt.Println(s)
}
