package main

import (
	"fmt"
)

func main() {
	// ================================================
	// soal 1
	fmt.Println("\n================= JAWABAN 1 =================")
	for i := 1; i <= 20; i++ {
		if i%2 == 1 {
			if i%3 == 0 {
				fmt.Printf("%d - I Love Coding\n", i)
			} else {
				fmt.Printf("%d - Santai\n", i)
			}
		} else {
			fmt.Printf("%d - Berkualitas\n", i)
		}
	}

	// ================================================
	// soal 2
	fmt.Println("\n================= JAWABAN 2 =================")
	for i := 0; i < 7; i++ {
		for j := 0; j < 7; j++ {
			if j <= i {
				fmt.Print("#")
			} else {
				break
			}
		}
		fmt.Println()
	}

	// ================================================
	// soal 3
	fmt.Println("\n================= JAWABAN 3 =================")
	var kalimat = [...]string{"aku", "dan", "saya", "sangat", "senang", "belajar", "golang"}
	fmt.Println(kalimat[2:])

	// ================================================
	// soal 4
	fmt.Println("\n================= JAWABAN 4 =================")
	var sayuran = []string {
		"Bayam",
		"Buncis",
		"Kangkung",
		"Kubis",
		"Seledri",
		"Tauge",
		"Timun",
	}

	for i, sayur := range sayuran {
		fmt.Printf("%d. %s\n", i+1, sayur)
	}

	// ================================================
	// soal 5
	fmt.Println("\n================= JAWABAN 5 =================")
	var satuan = map[string]int{
		"panjang": 7,
		"lebar":   4,
		"tinggi":  6,
	}

	var volumeBalok = 1
	for key, value := range satuan {
		fmt.Printf("%s = %d\n", key, value)
		volumeBalok *= value
	}
	fmt.Printf("Volume Balok = %d", volumeBalok)
}
