package main

import (
	"fmt"
	"unicode"
)

var jumlah = 0
var jumlahI = 0

func deret(i int) int {
	if i <= 1 {
		jumlah = jumlah + 1
		fmt.Print(1)
		return 1
	}
	y := deret(i-1) * 10
	jumlah = jumlah + y*i
	fmt.Printf(" + %d", y*i)
	return y
}

func pecahan(i int) int {
	if i <= 1 {
		jumlahI = jumlahI + 1
		jumlah = jumlah + 2
		fmt.Print("1/2x")
		return i
	}
	y := pecahan(i - 1)
	jumlahI = jumlahI + i
	x := i * 2
	jumlah = jumlah + x
	fmt.Printf(" %d/%dx", i, x)

	return y
}

func kapital(i int, x string) int {
	if i < 0 {
		return i
	}
	y := rune(x[i])
	if unicode.IsLower(y) {
		jumlah = jumlah + 1
	}

	return kapital(i-1, x)
}

func main() {
	var batas int
	var pilih int
	var kata string

	fmt.Println("1. Soal nomor 1")
	fmt.Println("2. Soal nomor 2")
	fmt.Println("3. Soal nomor 3")
	fmt.Print("Masukkan pilihan soal \t:")
	fmt.Scan(&pilih)

	switch pilih {
	case 1:
		fmt.Print("Masukkan batasan: ")
		fmt.Scan(&batas)

		deret(batas)
	case 2:
		fmt.Print("Masukkan batasan: ")
		fmt.Scan(&batas)

		pecahan(batas)
		fmt.Printf(" = %d/%dx", jumlahI, jumlah)
	case 3:
		fmt.Print("Masukkan sebuah kata: ")
		fmt.Scan(&kata)

		kapital(len(kata)-1, kata)
		fmt.Printf("jumlah huruf kecil: %d", jumlah)
	}
}
