package functions

import (
	"math"
	"strconv"
)

// ================================================
// Function soal 3
func LuasPersegi(value int, status bool) (text interface{}) {
	if status {
		if value == 0 {
			text = "Maaf anda belum menginput sisi dari persegi"
		} else {
			text = int32(math.Pow(float64(value), 2))
		}
	} else {
		if value == 0 {
			text = nil
		} else {
			text = value
		}
	}

	return
}

// ================================================
// Function soal 4
func ProcessOutput(s *string, total *int, arrPertama interface{}, arrKedua interface{}) {
	for i, angka := range arrPertama.([]int) {
		*total += angka
		*s += strconv.FormatInt(int64(angka), 10)

		if i != len(arrPertama.([]int))-1 {
			*s += " + "
		} else {
			if arrKedua != nil && len(arrKedua.([]int)) != 0 {

				*s += " + "
			} else {
				*s += " = " + strconv.FormatInt(int64(*total), 10)
			}
		}
	}
}
