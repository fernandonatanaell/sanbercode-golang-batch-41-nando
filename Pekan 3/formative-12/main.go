package main

import (
	"fmt"
	middleware "formative-12/pkg/middleware"
	models "formative-12/pkg/models"
	"net/http"
)

func main() {
	var nilaiNilaiMahasiswa = models.KartuHasilStudi{}

	// konfigurasi server
	server := &http.Server{
		Addr: ":8080",
	}
	// routing
	http.Handle("/", middleware.CekLogin(http.HandlerFunc(nilaiNilaiMahasiswa.DoNilaiMahasiswa)))

	// jalankan server
	fmt.Println("server running at http://localhost:8080")
	server.ListenAndServe()
}
