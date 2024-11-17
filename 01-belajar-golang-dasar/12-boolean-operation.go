package main

import "fmt"

func main() {
	// Boolean operator
	// &&, ||, !

	nilaiAkhir := 80
	absensi := 80

	lulusNilaiAkhir := nilaiAkhir >= 80
	lulusAbsensi := absensi >= 80

	lulus := lulusNilaiAkhir && lulusAbsensi

	fmt.Println(lulus)
}
