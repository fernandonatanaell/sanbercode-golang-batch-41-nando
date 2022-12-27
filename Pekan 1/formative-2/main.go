package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	// ================================================
	// soal 1
	var kataBootcamp = "Bootcamp"
	var kataDigital = "Digital"
	var kataSkill = "Skill"
	var kataSanbercode = "Sanbercode"
	var kataGolang = "Golang"

	fmt.Printf("%s %s %s %s %s \n", kataBootcamp, kataDigital, kataSkill, kataSanbercode, kataGolang)

	// ================================================
	// soal 2
	halo := "Halo Dunia"
	fmt.Println(strings.Replace(halo, "Dunia", "Golang", -1))

	// ================================================
	// soal 3
	var kataPertama = "saya"
	var kataKedua = "senang"
	var kataKetiga = "belajar"
	var kataKeempat = "golang"

	kataKedua = strings.Replace(kataKedua, string(kataKedua[0]), string(strings.ToUpper(kataKedua)[0]), 1)
	kataKetiga = kataKetiga[0:len(kataKetiga)-1] + string(strings.ToUpper(kataKetiga)[len(kataKetiga)-1])
	kataKeempat = strings.ToUpper(kataKeempat)

	fmt.Printf("%s %s %s %s \n", kataPertama, kataKedua, kataKetiga, kataKeempat)

	// ================================================
	// soal 4
	var angkaPertama = "8"
	var angkaKedua = "5"
	var angkaKetiga = "6"
	var angkaKeempat = "7"

	var numPertama, _ = strconv.ParseInt(angkaPertama, 10, 8)
	var numKedua, _ = strconv.ParseInt(angkaKedua, 10, 8)
	var numKetiga, _ = strconv.ParseInt(angkaKetiga, 10, 8)
	var numKeempat, _ = strconv.ParseInt(angkaKeempat, 10, 8)

	fmt.Println(numPertama + numKedua + numKetiga + numKeempat)

	// ================================================
	// soal 5
	kalimat := "halo halo bandung"
	angka := 2021

	kalimat = strings.ReplaceAll(kalimat, "halo", "Hi")
	fmt.Printf("\"%s\" - %d", kalimat, angka)
}
