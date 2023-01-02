package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// ================= START soal 1 =================
	fmt.Println("\n================= JAWABAN 1 =================")
	var luasLigkaran float64
	var kelilingLingkaran float64

	lingkaran(14, &luasLigkaran, &kelilingLingkaran)
	fmt.Printf("Luas lingkaran : %.2f cm2\n", luasLigkaran)
	fmt.Printf("Keliling lingkaran : %.2f cm\n", kelilingLingkaran)
	// ================= END soal 1 =================

	// ================= START soal 2 =================
	fmt.Println("\n================= JAWABAN 2 =================")
	var sentence string
	introduce(&sentence, "John", "laki-laki", "penulis", "30")
	fmt.Println(sentence)

	introduce(&sentence, "Sarah", "perempuan", "model", "28")
	fmt.Println(sentence)
	// ================= END soal 2 =================

	// ================= START soal 3 =================
	fmt.Println("\n================= JAWABAN 3 =================")
	var buah = []string{}
	addBuah(&buah, "Jeruk")
	addBuah(&buah, "Semangka")
	addBuah(&buah, "Mangga")
	addBuah(&buah, "Strawberry")
	addBuah(&buah, "Durian")
	addBuah(&buah, "Manggis")
	addBuah(&buah, "Alpukat")

	for i, item := range buah {
		fmt.Printf("%d. %s\n", i+1, item)
	}
	// ================= END soal 3 =================

	// ================= START soal 4 =================
	fmt.Println("\n================= JAWABAN 4 =================")
	var dataFilm = []map[string]string{}
	tambahDataFilm("LOTR", "2 jam", "action", "1999", &dataFilm)
	tambahDataFilm("avenger", "2 jam", "action", "2019", &dataFilm)
	tambahDataFilm("spiderman", "2 jam", "action", "2004", &dataFilm)
	tambahDataFilm("juon", "2 jam", "horror", "2004", &dataFilm)

	for i, item := range dataFilm {
		var tempString string

		if item["title"] == "spiderman" {
			tempString = strconv.FormatInt(int64(i+1), 10) + ". genre : " + item["genre"] +
				"\n   year : " + item["tahun"] +
				"\n   title : " + item["title"] +
				"\n   duration : " + item["jam"] + "\n"
		} else {
			tempString = strconv.FormatInt(int64(i+1), 10) + ". title : " + item["title"] +
				"\n   duration : " + item["jam"] +
				"\n   genre : " + item["genre"] +
				"\n   year : " + item["tahun"] + "\n"
		}

		fmt.Println(tempString)
	}
	// ================= END soal 4 =================
}

// ================================================
// function soal 1
func lingkaran(jari float64, luasLingkaran *float64, kelilingLingkaran *float64) {
	*luasLingkaran = 3.14 * jari * jari
	*kelilingLingkaran = 3.14 * 2 * jari
}

// ================================================
// function soal 2
func introduce(sentence *string, nama string, jk string, profesi string, umur string) {
	if strings.ToLower(jk) == "laki-laki" {
		*sentence = "Pak "
	} else {
		*sentence = "Bu "
	}

	*sentence += nama + " adalah seorang " + profesi + " yang berusia " + umur + " tahun"
}

// ================================================
// function soal 3
func addBuah(buah *[]string, buahName string) {
	*buah = append(*buah, buahName)
}

// ================================================
// function soal 4
func tambahDataFilm(judul string, durasi string, genre string, tahun string, dataFilm *[]map[string]string) {
	*dataFilm = append(*dataFilm, map[string]string{
		"genre": genre,
		"jam":   durasi,
		"tahun": tahun,
		"title": judul,
	})
}
