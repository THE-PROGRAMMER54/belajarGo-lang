package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Mahasiswa struct {
	Nim    int
	Nama   string
	Alamat string
	Umur   int
	Nilai  float64
}

var dataMahasiswa = make(map[int]Mahasiswa)
var reader = bufio.NewReader(os.Stdin)

func cruddenganMap() {
	for {
		fmt.Println("\n=== Menu ===")
		fmt.Println("1. Tambah Data")
		fmt.Println("2. Tampilkan Data")
		fmt.Println("3. Ubah Data")
		fmt.Println("4. Hapus Data")
		fmt.Println("5. Keluar")

		fmt.Print("Pilih Menu 1-5:")
		var pilihan int
		fmt.Scanln(&pilihan)

		switch pilihan {
		case 1:
			tambahData()
		case 2:
			tampilkanData()
		case 3:
			ubahData()
		case 4:
			hapusData()
		case 5:
			fmt.Println("Terim Kasih..")
			return
		default:
			fmt.Println("❌ Pilihan tidak valid.")
		}
	}
}

func input(prompt string) string {
	fmt.Print(prompt)
	text, _ := reader.ReadString('\n')
	text = strings.TrimSpace(text)
	return text
}

func tambahData() {
	nimStr := input("Masukan nim:")
	nim, err := strconv.Atoi(nimStr)
	if err != nil {
		fmt.Println("⚠️  Nim harus berupa angka.")
		return
	}

	if _, exists := dataMahasiswa[nim]; exists {
		fmt.Println("⚠️  NIM sudah terdaftar.")
		return
	}

	nama := input("Masukan Nama Anda:")
	alamat := input("Masukan Alamat Anda:")
	umurStr := input("Masukan Umur Anda:")
	nilaiStr := input("Masukan Nilai Anda:")

	umur, err1 := strconv.Atoi(umurStr)
	nilai, err2 := strconv.ParseFloat(nilaiStr, 64)

	if err1 != nil || err2 != nil {
		fmt.Println("⚠️  Umur dan nilai harus berupa angka.")
		return
	}

	dataMahasiswa[nim] = Mahasiswa{
		Nim:    nim,
		Nama:   nama,
		Alamat: alamat,
		Umur:   umur,
		Nilai:  nilai,
	}

	fmt.Println("✅ Data berhasil ditambahkan.")
}

func tampilkanData() {
	fmt.Println("\n=== Data Mahasiswa ===")

	if len(dataMahasiswa) == 0 {
		fmt.Println("Belum ada data.")
		return
	}

	for nim, mhs := range dataMahasiswa {
		fmt.Printf("NIM: %d | Nama: %s | Umur: %d | Nilai: %.1f\n", nim, mhs.Nama, mhs.Umur, mhs.Nilai)
	}
}

func ubahData() {

}

func hapusData() {

}
