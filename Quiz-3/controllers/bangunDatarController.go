package controllers

import (
	"fmt"
	"math"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"
)

func SegitigaSamaSisi(ctx *gin.Context) {
	hitung := strings.ToLower(ctx.Query("hitung"))
	alas, errAlas := strconv.Atoi(ctx.Query("alas"))
	tinggi, errTinggi := strconv.Atoi(ctx.Query("tinggi"))
	var result string

	if hitung != "luas" && hitung != "keliling" {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Parameter hitung invalid!",
			"error_message": `Hitung harus diisi "luas" atau "keliling" saja!`,
		})
		return
	}

	if errAlas != nil || errTinggi != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Parameter alas/tinggi invalid!",
			"error_message": "Alas/tinggi harus berupa angka!",
		})
		return
	}

	if hitung == "luas" {
		result = fmt.Sprintf("Luas segitiga sama sisi: %d cm2", (1*alas*tinggi)/2)
	} else {
		result = fmt.Sprintf("Keliling segitiga sama sisi: %d cm", alas*3)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}

func Persegi(ctx *gin.Context) {
	hitung := strings.ToLower(ctx.Query("hitung"))
	sisi, errSisi := strconv.Atoi(ctx.Query("sisi"))
	var result string

	if hitung != "luas" && hitung != "keliling" {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Parameter hitung invalid!",
			"error_message": `Hitung harus diisi "luas" atau "keliling" saja!`,
		})
		return
	}

	if errSisi != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Parameter sisi invalid!",
			"error_message": "Sisi harus berupa angka!",
		})
		return
	}

	if hitung == "luas" {
		result = fmt.Sprintf("Luas persegi: %d cm2", sisi*sisi)
	} else {
		result = fmt.Sprintf("Keliling persegi: %d cm", sisi*4)
	}

	ctx.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}

func PersegiPanjang(ctx *gin.Context) {
	hitung := strings.ToLower(ctx.Query("hitung"))
	panjang, errPanjang := strconv.Atoi(ctx.Query("panjang"))
	lebar, errLebar := strconv.Atoi(ctx.Query("lebar"))
	var result string

	if hitung != "luas" && hitung != "keliling" {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Parameter hitung invalid!",
			"error_message": `Hitung harus diisi "luas" atau "keliling" saja!`,
		})
		return
	}

	if errPanjang != nil || errLebar != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Parameter panjang/lebar invalid!",
			"error_message": "Panjang/lebar harus berupa angka!",
		})
		return
	}

	if hitung == "luas" {
		result = fmt.Sprintf("Luas persegi panjang: %d cm2", panjang*lebar)
	} else {
		result = fmt.Sprintf("Keliling persegi panjang: %d cm", ((2 * panjang) + (2 * lebar)))
	}

	ctx.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}

func Lingkaran(ctx *gin.Context) {
	hitung := strings.ToLower(ctx.Query("hitung"))
	jariJari, errJariJari := strconv.Atoi(ctx.Query("jariJari"))
	var result string

	if hitung != "luas" && hitung != "keliling" {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Parameter hitung invalid!",
			"error_message": `Hitung harus diisi "luas" atau "keliling" saja!`,
		})
		return
	}

	if errJariJari != nil {
		ctx.AbortWithStatusJSON(http.StatusNotFound, gin.H{
			"error_status":  "Parameter jari-jari invalid!",
			"error_message": "Jari-jari harus berupa angka!",
		})
		return
	}

	if hitung == "luas" {
		result = fmt.Sprintf("Luas lingkaran: %.1f cm2", math.Pi*float32(jariJari)*float32(jariJari))
	} else {
		result = fmt.Sprintf("Keliling lingkaran: %.1f cm", math.Pi*(2 * float32(jariJari)))
	}

	ctx.JSON(http.StatusOK, gin.H{
		"result": result,
	})
}
