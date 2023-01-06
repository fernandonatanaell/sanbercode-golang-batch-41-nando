package functions

import (
	"errors"
	"fmt"
	"sort"
	"strconv"
	"sync"
	"time"
)

// ================================================
// Function soal 1
func CetakKalimat(kalimat string, tahun int) {
	fmt.Println("\n================= JAWABAN 1 =================")
	fmt.Printf("%s %d \n\n", kalimat, tahun)
}

// ================================================
// Function soal 2
func KelilingSegitigaSamaSisi(sisi int, status bool) (messageSuccess string, messageError error) {
	defer restoreKeliling(&messageError)

	if status {
		if sisi == 0 {
			panic("Maaf anda belum menginput sisi dari segitiga sama sisi")
		} else {
			messageSuccess = " keliling segitiga sama sisinya dengan sisi " + strconv.Itoa(sisi) + " cm adalah " + strconv.Itoa(sisi*3) + " cm"
		}
	} else {
		if sisi == 0 {
			panic("Maaf anda belum menginput sisi dari segitiga sama sisi")
		} else {
			messageSuccess = " " + strconv.Itoa(sisi)
		}
	}

	return
}

func restoreKeliling(err *error) {
	message := recover()

	if message == nil {
		*err = errors.New("")
	} else {
		*err = errors.New(fmt.Sprintf("%v", message))
	}
}

// ================================================
// Function soal 3
func TambahAngka(new_number int, old_number *int) {
	*old_number += new_number
}

func CetakAngka(number *int) {
	fmt.Println("\n================= JAWABAN 3 =================")
	fmt.Println("Total:", *number)
}

// ================================================
// Function soal 4
func TambahPhone(name_phone string, phone *[]string) {
	*phone = append(*phone, name_phone)
	sort.Strings(*phone)
}

// ================================================
// Function soal 5
func CetakPhones() {
	var wg sync.WaitGroup
	defer wg.Wait()

	var phones = []string{"Xiaomi", "Asus", "Iphone", "Samsung", "Oppo", "Realme", "Vivo"}
	sort.Strings(phones)

	for i, phone := range phones {
		wg.Add(1)
		go cetakPhone(i, phone, &wg)
		time.Sleep(1 * time.Second)
	}
}

func cetakPhone(index int, name_phone string, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("%d. %s\n", index+1, name_phone)
}

// ================================================
// Function soal 6
func GetMovies(moviesChannel chan string, movies ...string) {
	for _, movie := range movies {
		moviesChannel <- movie
	}

	close(moviesChannel)
}
