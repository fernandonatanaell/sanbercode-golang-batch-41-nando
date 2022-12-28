package main

import (
	"fmt"
	"strconv"
)

func main() {
	// ================================================
	// soal 1
	fmt.Println("\n================= JAWABAN 1 =================")
	var panjangPersegiPanjang string = "8"
	var lebarPersegiPanjang string = "5"
	var numPanjangPersegiPanjang, _ = strconv.ParseInt(panjangPersegiPanjang, 10, 8)
	var numLebarPersegiPanjang, _ = strconv.ParseInt(lebarPersegiPanjang, 10, 8)

	var alasSegitiga string = "6"
	var tinggiSegitiga string = "7"
	var numAlasSegitiga, _ = strconv.ParseInt(alasSegitiga, 10, 8)
	var numTinggiSegitiga, _ = strconv.ParseInt(tinggiSegitiga, 10, 8)

	var luasPersegiPanjang int = int(numPanjangPersegiPanjang) * int(numLebarPersegiPanjang)
	var kelilingPersegiPanjang int = 2 * (int(numPanjangPersegiPanjang) + int(numLebarPersegiPanjang))
	var luasSegitiga int = (1 * int(numAlasSegitiga) * int(numTinggiSegitiga)) / 2

	fmt.Printf("Luas Persegi Panjang : %d cm2 \n", luasPersegiPanjang)
	fmt.Printf("Keliling Persegi Panjang : %d cm \n", kelilingPersegiPanjang)
	fmt.Printf("Luas Segitiga : %d cm2 \n\n", luasSegitiga)

	// ================================================
	// soal 2
	fmt.Println("================= JAWABAN 2 =================")
	var nilaiJohn = 80
	var nilaiDoe = 50

	if nilaiJohn >= 80 {
		fmt.Println("indeksnya A")
	} else if nilaiJohn >= 70 && nilaiJohn < 80 {
		fmt.Println("indeksnya B")
	} else if nilaiJohn >= 60 && nilaiJohn < 70 {
		fmt.Println("indeksnya C")
	} else if nilaiJohn >= 50 && nilaiJohn < 60 {
		fmt.Println("indeksnya D")
	} else {
		fmt.Println("indeksnya E")
	}

	if nilaiDoe >= 80 {
		fmt.Println("indeksnya A")
	} else if nilaiDoe >= 70 && nilaiDoe < 80 {
		fmt.Println("indeksnya B")
	} else if nilaiDoe >= 60 && nilaiDoe < 70 {
		fmt.Println("indeksnya C")
	} else if nilaiDoe >= 50 && nilaiDoe < 60 {
		fmt.Println("indeksnya D")
	} else {
		fmt.Println("indeksnya E")
	}

	// ================================================
	// soal 3
	fmt.Println("\n================= JAWABAN 3 =================")
	var tanggal = 18
	var bulan = 12
	var tahun = 2002

	var namaBulan string
	switch bulan {
	case 1:
		namaBulan = "Januari"
	case 2:
		namaBulan = "Februari"
	case 3:
		namaBulan = "Maret"
	case 4:
		namaBulan = "April"
	case 5:
		namaBulan = "Mei"
	case 6:
		namaBulan = "Juni"
	case 7:
		namaBulan = "Juli"
	case 8:
		namaBulan = "Agustus"
	case 9:
		namaBulan = "September"
	case 10:
		namaBulan = "Oktober"
	case 11:
		namaBulan = "November"
	case 12:
		namaBulan = "Desember"
	default:
		namaBulan = "-"
	}

	if namaBulan == "-" {
		fmt.Println("Tanggal invalid!")
	} else {
		fmt.Printf("%d %s %d \n", tanggal, namaBulan, tahun)
	}

	// ================================================
	// soal 4
	fmt.Println("\n================= JAWABAN 4 =================")
	if tahun >= 1944 && tahun <= 1964 {
		fmt.Println("Baby boomer")
	} else if tahun >= 1965 && tahun <= 1979 {
		fmt.Println("Generasi X")
	} else if tahun >= 1980 && tahun <= 1994 {
		fmt.Println("Generasi Y (Millenials)")
	} else if tahun >= 1995 && tahun <= 2015 {
		fmt.Println("Generasi Z")
	}
}
