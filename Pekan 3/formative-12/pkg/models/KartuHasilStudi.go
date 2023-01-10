package models

import (
	"encoding/json"
	"net/http"
	"strconv"
)

// Struct
type KartuHasilStudi struct {
	NilaiNilaiMahasiswa []NilaiMahasiswa
}

// Function yang akan dijalnkan ketika localhost:8080/ diakses
func (khs *KartuHasilStudi) DoNilaiMahasiswa(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" { // Method POST
		var newKHS NilaiMahasiswa
		var tempNilai uint32

		if r.Header.Get("Content-Type") == "application/json" {
			// parse dari json
			decodeJSON := json.NewDecoder(r.Body)
			if err := decodeJSON.Decode(&newKHS); err != nil {
				http.Error(w, err.Error(), http.StatusNotFound)
				return
			}
			tempNilai = newKHS.Nilai
		} else {
			// parse dari form
			nama := r.PostFormValue("Nama")
			mataKuliah := r.PostFormValue("MataKuliah")
			getNilai := r.PostFormValue("Nilai")
			nilai, errNilai := strconv.Atoi(getNilai)

			// CHECK APAKAH NILAI YANG DIPASSINGKAN INTEGER ATAU BUKAN
			if errNilai != nil {
				http.Error(w, "Input nilai invalid!", http.StatusNotFound)
				return
			}

			tempNilai = uint32(nilai)
			newKHS = NilaiMahasiswa{
				Nama:       nama,
				MataKuliah: mataKuliah,
				Nilai:      uint32(nilai),
			}
		}

		// CHECK APAKAH INPUT ADA YANG KOSONG ATAU TIDAK?
		if !checkInput(newKHS.Nama, newKHS.MataKuliah, newKHS.Nilai) {
			http.Error(w, "Input nama atau mata kuliah atau nilai INVALID", http.StatusNotFound)
			return
		}

		// CHECK APAKAH NILAI YANG DIINPUTKAN LEBIH DARI 100?
		if tempNilai > 100 {
			http.Error(w, "Nilai tidak boleh > dari 100", http.StatusNotFound)
			return
		}

		// INPUT AMAN SEMUA, TAMBAHKAN KE DALAM SLICE
		newKHS.IndeksNilai = getIndeksPrestasi(tempNilai)
		newKHS.ID = getLastID(khs.NilaiNilaiMahasiswa)
		khs.NilaiNilaiMahasiswa = append(khs.NilaiNilaiMahasiswa, newKHS)

		w.Write([]byte("Nilai Mahasiswa berhasil dimasukkan!")) // cetak di browser
	} else { // Method GET
		dataKHS, _ := json.Marshal(khs.NilaiNilaiMahasiswa) // to byte
		w.Write(dataKHS)                                    // cetak di browser
	}

	return
}

func getIndeksPrestasi(nilai uint32) string {
	if nilai >= 80 {
		return "A"
	} else if nilai >= 70 && nilai < 80 {
		return "B"
	} else if nilai >= 60 && nilai < 70 {
		return "C"
	} else if nilai >= 50 && nilai < 60 {
		return "D"
	} else {
		return "E"
	}
}

func getLastID(nm []NilaiMahasiswa) uint32 {
	if lengthArr := len(nm); lengthArr == 0 {
		return 1
	} else {
		return uint32(nm[len(nm)-1].ID + 1)
	}
}

func checkInput(nama string, matkul string, nilai uint32) bool {
	return !(nama == "" || matkul == "" || nilai == 0)
}
