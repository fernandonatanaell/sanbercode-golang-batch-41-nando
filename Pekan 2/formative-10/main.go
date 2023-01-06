package main

import (
	"fmt"
	functions "formative-10/pkg/functions"
	"time"
)

func main() {
	// deklarasi variabel angka ini simpan di baris pertama func main
	angka := 1

	// ================= SOAL 1 =================
	defer functions.CetakKalimat("Golang Backend Development", 2021)

	// ================= SOAL 2 =================
	fmt.Println("\n================= JAWABAN 2 =================")
	fmt.Println(functions.KelilingSegitigaSamaSisi(4, true))
	fmt.Println(functions.KelilingSegitigaSamaSisi(8, false))
	fmt.Println(functions.KelilingSegitigaSamaSisi(0, true))
	fmt.Println(functions.KelilingSegitigaSamaSisi(0, false))

	// ================= SOAL 3 =================
	defer functions.CetakAngka(&angka)

	functions.TambahAngka(7, &angka)
	functions.TambahAngka(6, &angka)
	functions.TambahAngka(-1, &angka)
	functions.TambahAngka(9, &angka)

	// ================= SOAL 4 =================
	fmt.Println("\n================= JAWABAN 4 =================")
	var phones = []string{}

	functions.TambahPhone("Xiaomi", &phones)
	functions.TambahPhone("Asus", &phones)
	functions.TambahPhone("IPhone", &phones)
	functions.TambahPhone("Samsung", &phones)
	functions.TambahPhone("Oppo", &phones)
	functions.TambahPhone("Realme", &phones)
	functions.TambahPhone("Vivo", &phones)

	for i, phone := range phones {
		fmt.Printf("%d. %s\n", i+1, phone)
		time.Sleep(1 * time.Second)
	}

	// ================= SOAL 5 =================
	fmt.Println("\n================= JAWABAN 5 =================")
	functions.CetakPhones()

	// ================= SOAL 6 =================
	fmt.Println("\n================= JAWABAN 6 =================")
	var movies = []string{"Harry Potter", "LOTR", "SpiderMan", "Logan", "Avengers", "Insidious", "Toy Story"}
	moviesChannel := make(chan string)

	go functions.GetMovies(moviesChannel, movies...)

	for value := range moviesChannel {
		fmt.Println(value)
	}
}
