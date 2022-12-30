package main

import (
	"fmt"
	"strings"
)

func main() {
	// ================================================
	// soal 1
	fmt.Println("\n================= JAWABAN 1 =================")
	panjang := 12
	lebar := 4
	tinggi := 8

	luas := luasPersegiPanjang(panjang, lebar)
	keliling := kelilingPersegiPanjang(panjang, lebar)
	volume := volumeBalok(panjang, lebar, tinggi)

	fmt.Println(luas)
	fmt.Println(keliling)
	fmt.Println(volume)

	// ================================================
	// soal 2
	fmt.Println("\n================= JAWABAN 2 =================")
	john := introduce("John", "laki-laki", "penulis", "30")
	fmt.Println(john) // Menampilkan "Pak John adalah seorang penulis yang berusia 30 tahun"

	sarah := introduce("Sarah", "perempuan", "model", "28")
	fmt.Println(sarah) // Menampilkan "Bu Sarah adalah seorang model yang berusia 28 tahun"

	// ================================================
	// soal 3
	fmt.Println("\n================= JAWABAN 3 =================")
	var buah = []string{"semangka", "jeruk", "melon", "pepaya"}
	var buahFavoritJohn = buahFavorit("John", buah...)

	fmt.Println(buahFavoritJohn)
	// halo nama saya john dan buah favorit saya adalah "semangka", "jeruk", "melon", "pepaya"

	// ================================================
	// soal 4
	fmt.Println("\n================= JAWABAN 4 =================")
	var dataFilm = []map[string]string{}
	// buatlah closure function disini
	var tambahDataFilm = func(film ...string) {
		if len(film) == 4 {
			dataFilm = append(dataFilm, map[string]string{
				"genre": film[2],
				"jam":   film[1],
				"tahun": film[3],
				"title": film[0],
			})
		}
	}

	tambahDataFilm("LOTR", "2 jam", "action", "1999")
	tambahDataFilm("avenger", "2 jam", "action", "2019")
	tambahDataFilm("spiderman", "2 jam", "action", "2004")
	tambahDataFilm("juon", "2 jam", "horror", "2004")

	for _, item := range dataFilm {
		fmt.Println(item)
	}
}

// ================================================
// function soal 1
func luasPersegiPanjang(panjang int, lebar int) int {
	return panjang * lebar
}
func kelilingPersegiPanjang(panjang int, lebar int) int {
	return 2 * (panjang + lebar)
}
func volumeBalok(panjang int, lebar int, tinggi int) int {
	return panjang * lebar * tinggi
}

// ================================================
// function soal 2
func introduce(nama string, jk string, profesi string, umur string) (text string) {
	if strings.ToLower(jk) == "laki-laki" {
		text = "Pak "
	} else {
		text = "Bu "
	}

	text += nama + " adalah seorang " + profesi + " yang berusia " + umur + " tahun"
	return
}

// ================================================
// function soal 3
func buahFavorit(nama string, buahs ...string) string {
	var text string = "halo nama saya " + strings.ToLower(nama) + " dan buah favorit saya adalah "
	for i, buah := range buahs {
		text += "\"" + buah + "\""

		if i != len(buahs)-1 {
			text += ", "
		}
	}
	return text
}
