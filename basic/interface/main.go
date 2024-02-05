package main

import "fmt"


//kontrak  aturan
type BangunDatar interface {
	HitungLuas() int
}

// implementasi dari aturan
type Persegi struct {
	Sisi int
}

func (persegi Persegi) HitungLuas() int{
	return persegi.Sisi * persegi.Sisi
}

type PersegiPanjang struct {
	Panjang int
	Lebar int 
}

func (persegiPanjang PersegiPanjang) HitungLuas() int{
	return persegiPanjang.Panjang * persegiPanjang.Lebar
}

func main() {
	persegiPanjang := PersegiPanjang{Panjang: 4, Lebar: 8}
	luas := SeberapaLuas(persegiPanjang)
	fmt.Println("Luas Persegi Panjang: ", luas)

	persegi := Persegi{Sisi: 8}
	luas = SeberapaLuas(persegi)
	fmt.Println("Luas Persegi: ", luas)

}

func SeberapaLuas(bangunDatar BangunDatar) int {
	return bangunDatar.HitungLuas()
} 


