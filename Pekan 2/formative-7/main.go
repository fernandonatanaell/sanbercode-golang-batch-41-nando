package main

import (
	"fmt"
	"strconv"
	"strings"
)

// ================================================
// Struct soal 1
type buah struct {
	nama, warna string
	adaBijinya  bool
	harga       int32
}

// ================================================
// Struct & method soal 2
type segitiga struct {
	alas, tinggi int
}

func (s segitiga) luas() float32 {
	return 0.5 * float32(s.alas) * float32(s.tinggi)
}

type persegi struct {
	sisi int
}

func (p persegi) luas() int {
	return p.sisi * p.sisi
}

type persegiPanjang struct {
	panjang, lebar int
}

func (pp persegiPanjang) luas() int {
	return pp.panjang * pp.lebar
}

// ================================================
// Struct & method soal 3
type phone struct {
	name, brand string
	year        int
	colors      []string
}

func (p *phone) addColor(new_color string) {
	p.colors = append(p.colors, new_color)
}

// ================================================
// Struct & function soal 4
type movie struct {
	title, genre   string
	duration, year int
}

func tambahDataFilm(judul string, durasi int, genre string, tahun int, dataFilm *[]movie) {
	*dataFilm = append(*dataFilm, movie{
		genre:    genre,
		duration: durasi,
		year:     tahun,
		title:    judul,
	})
}

func minutesToHours(minutes int) string {
	var tempHour float32 = float32(minutes) / 60
	s := fmt.Sprintf("%.1f", tempHour)

	return strings.TrimRight(strings.TrimRight(s, "0"), ".") + " jam"
}

func main() {
	// ================= START soal 1 =================
	fmt.Println("\n================= JAWABAN 1 =================")
	var buahs = []buah{
		{nama: "Nanas", warna: "Kuning", adaBijinya: false, harga: 9000},
		{nama: "Jeruk", warna: "Oranye", adaBijinya: true, harga: 8000},
		{nama: "Semangka", warna: "Hijau & Merah", adaBijinya: true, harga: 10000},
		{nama: "Pisang", warna: "Kuning", adaBijinya: false, harga: 5000},
	}

	fmt.Printf("%-15s %-15s %-14s %-6s \n", "Nama", "Warna", "Ada Bijinya", "Harga")
	for _, item := range buahs {
		var statusBiji = "Ada"
		if !item.adaBijinya {
			statusBiji = "Tidak"
		}

		fmt.Printf("%-15s %-15s %-14s %-6d \n", item.nama, item.warna, statusBiji, item.harga)
	}
	// ================= END soal 1 =================

	// ================= START soal 2 =================
	fmt.Println("\n================= JAWABAN 2 =================")
	var segitiga = segitiga{alas: 8, tinggi: 7}
	fmt.Printf("Luas segitiga: %.2f cm\n", segitiga.luas())

	var persegi = persegi{sisi: 6}
	fmt.Printf("Luas persegi: %d cm\n", persegi.luas())

	var persegiPanjang = persegiPanjang{panjang: 15, lebar: 4}
	fmt.Printf("Luas persegi panjang: %d cm\n", persegiPanjang.luas())
	// ================= END soal 2 =================

	// ================= START soal 3 =================
	fmt.Println("\n================= JAWABAN 3 =================")
	var phones = []phone{
		{name: "Iphone 13", brand: "Iphone", year: 2001, colors: []string{}},
		{name: "Samsung S20 FE", brand: "Samsung", year: 2002, colors: []string{}},
		{name: "Vivo F11", brand: "Vivo", year: 2003, colors: []string{}},
	}

	phones[0].addColor("Merah")
	phones[0].addColor("Kuning")

	phones[1].addColor("Hijau")
	phones[1].addColor("Putih")

	phones[2].addColor("Hitam")
	phones[2].addColor("Ungu")

	fmt.Println(phones)
	// ================= END soal 3 =================

	// ================= START soal 4 =================
	fmt.Println("\n================= JAWABAN 4 =================")
	var dataFilm = []movie{}

	tambahDataFilm("LOTR", 120, "action", 1999, &dataFilm)
	tambahDataFilm("avenger", 120, "action", 2019, &dataFilm)
	tambahDataFilm("spiderman", 120, "action", 2004, &dataFilm)
	tambahDataFilm("juon", 120, "horror", 2004, &dataFilm)

	for i, item := range dataFilm {
		var tempString string

		if item.title == "spiderman" {
			tempString = strconv.FormatInt(int64(i+1), 10) + ". genre : " + item.genre +
				"\n   year : " + strconv.Itoa(item.year) +
				"\n   title : " + item.title +
				"\n   duration : " + minutesToHours(item.duration) + "\n"
		} else {
			tempString = strconv.FormatInt(int64(i+1), 10) + ". title : " + item.title +
				"\n   duration : " + minutesToHours(item.duration) +
				"\n   genre : " + item.genre +
				"\n   year : " + strconv.Itoa(item.year) + "\n"
		}

		fmt.Println(tempString)
	}
	// ================= END soal 4 =================
}
