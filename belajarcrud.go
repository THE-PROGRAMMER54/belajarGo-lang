package main

// type buah struct {
// 	Nama   string
// 	Jumlah int
// 	Harga  float64
// }

// var Daftar []buah

// func crud() {
// 	for {
// 		fmt.Println("\n=== MENU ===")
// 		fmt.Println("1.Tambah Buah")
// 		fmt.Println("2.Daftar Buah")
// 		fmt.Println("3.Edit Buah")
// 		fmt.Println("4.Hapus Buah")
// 		fmt.Println("5.Keluar..")

// 		fmt.Print("Pilih Menu:")
// 		var pilihan int
// 		fmt.Scanln(&pilihan)

// 		switch pilihan {
// 		case 1:
// 			tambahBuah()
// 		case 2:
// 			daftarBuah()
// 		case 3:
// 			editBuah()
// 		case 4:
// 			hapusBuah()
// 		case 5:
// 			fmt.Println("Keluar ..")
// 			return
// 		}
// 	}
// }

// func tambahBuah() {
// 	reader := bufio.NewReader(os.Stdin)

// 	fmt.Print("Masukan nama buah: ")
// 	newbuah, _ := reader.ReadString('\n')
// 	newbuah = strings.TrimSpace(newbuah)
// 	if newbuah == "" {
// 		tambahBuah()
// 	}

// 	fmt.Print("Masukan jumlah buah: ")
// 	jumlahbuah, _ := reader.ReadString('\n')
// 	jumlahbuah = strings.TrimSpace(jumlahbuah)
// 	j, err := strconv.Atoi(jumlahbuah)
// 	if err != nil {
// 		fmt.Println("⚠️ Nilai harus berupa angka!")
// 		return
// 	}

// 	fmt.Print("Masukan harga buah: ")
// 	hargabuah, _ := reader.ReadString('\n')
// 	hargabuah = strings.TrimSpace(hargabuah)
// 	h, err := strconv.ParseFloat(hargabuah, 64)

// 	if err != nil {
// 		fmt.Println("⚠️ Nilai harus berupa angka!")
// 		return
// 	}

// 	Daftar = append(Daftar, buah{newbuah, j, h})

// 	fmt.Println("✅ Buah Berhasil Ditambahkan")
// }

// func daftarBuah() {
// 	fmt.Println("\n ===Macam-Macam Buah===")
// 	if len(Daftar) == 0 {
// 		fmt.Println("Belum ada buah.")
// 		return
// 	}

// 	for i, b := range Daftar {
// 		fmt.Printf("%d.Nama Buah: %s|Jumlah Buah %d|Harga Buah:%.0f\n", i+1, b.Nama, b.Jumlah, b.Harga)
// 	}
// }

// func editBuah() {
// 	fmt.Println("\n ===Daftar Buah===")
// 	if len(Daftar) == 0 {
// 		fmt.Println("Belum ada buah.")
// 		return
// 	}

// 	for i, b := range Daftar {
// 		fmt.Printf("%d.Nama Buah: %s|Jumlah Buah %d|Harga Buah:%.0f\n", i+1, b.Nama, b.Jumlah, b.Harga)
// 	}

// 	reader := bufio.NewReader(os.Stdin)
// 	fmt.Print("Masukan nama buah yang ingin di edit:")
// 	buah, _ := reader.ReadString('\n')
// 	buah = strings.TrimSpace(buah)

// 	found := false
// 	for i := range Daftar {
// 		if strings.EqualFold(Daftar[i].Nama, buah) {
// 			found = true
// 			fmt.Print("Masukan nama buah (baru): ")
// 			namabuah, _ := reader.ReadString('\n')
// 			namabuah = strings.TrimSpace(namabuah)
// 			if namabuah == "" {
// 				tambahBuah()
// 			}

// 			fmt.Print("Masukan jumlah buah (baru): ")
// 			jumlahbuah, _ := reader.ReadString('\n')
// 			jumlahbuah = strings.TrimSpace(jumlahbuah)
// 			j, err := strconv.Atoi(jumlahbuah)
// 			if err != nil {
// 				fmt.Println("⚠️ Nilai harus berupa angka!")
// 				return
// 			}

// 			fmt.Print("Masukan harga buah (baru): ")
// 			hargabuah, _ := reader.ReadString('\n')
// 			hargabuah = strings.TrimSpace(hargabuah)
// 			h, err := strconv.ParseFloat(hargabuah, 64)

// 			if err != nil {
// 				fmt.Println("⚠️ Nilai harus berupa angka!")
// 				return
// 			}

// 			Daftar[i].Nama = namabuah
// 			Daftar[i].Jumlah = j
// 			Daftar[i].Harga = h
// 		}
// 	}
// 	if !found {
// 		fmt.Println("❌ Nama buah tidak ditemukan.")
// 	}
// }

// func hapusBuah() {
// 	fmt.Println("\n ===Daftar Buah===")
// 	if len(Daftar) == 0 {
// 		fmt.Println("Belum Ada Buah")
// 		return
// 	}
// 	for i, b := range Daftar {
// 		fmt.Printf("%d.Nama Buah: %s|Jumlah Buah %d|Harga Buah:%.0f\n", i+1, b.Nama, b.Jumlah, b.Harga)
// 	}

// 	reader := bufio.NewReader(os.Stdin)
// 	fmt.Print("Masukan nama buah yang ingin di edit:")
// 	buah, _ := reader.ReadString('\n')
// 	buah = strings.TrimSpace(buah)

// 	found := false
// 	for i := range Daftar {
// 		if strings.EqualFold(Daftar[i].Nama, buah) {
// 			if i == len(Daftar)-1 {
// 				Daftar = Daftar[:i]
// 			} else {
// 				Daftar = append(Daftar[:i], Daftar[i+1])
// 			}

// 			fmt.Println("✅ Mahasiswa berhasil dihapus.")
// 			found = true
// 			break
// 		}
// 	}

// 	if !found {
// 		fmt.Println("❌ Nama tidak ditemukan.")
// 	}
// }
