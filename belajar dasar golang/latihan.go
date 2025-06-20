package main

// func latihan1() {
// 	reader := bufio.NewReader(os.Stdin)
// 	var nama, hobi string
// 	var umur int

// 	for {
// 		fmt.Print("Masukan nama:")
// 		namainput, _ := reader.ReadString('\n')
// 		nama = strings.TrimSpace(namainput)
// 		if nama == "" {
// 			fmt.Println("⚠️ Nama tidak boleh kosong!")
// 			continue
// 		}
// 		fmt.Print("Masukan hobi:")
// 		hobies, _ := reader.ReadString('\n')
// 		hobi = strings.TrimSpace(hobies)
// 		if hobi == "" {
// 			fmt.Println("⚠️ Hobi tidak boleh kosong!")
// 			continue
// 		}

// 		fmt.Print("Masukan umur:")
// 		umurstring, _ := reader.ReadString('\n')
// 		umurstring = strings.TrimSpace(umurstring)
// 		u, err := strconv.Atoi(umurstring)
// 		if err != nil {
// 			fmt.Println("⚠️ Umur harus berupa angka!")
// 			continue
// 		}

// 		umur = u

// 		if umur < 0 {
// 			fmt.Println("⚠️ Umur tidak boleh kosong!")
// 			continue
// 		} else if umur > 120 {
// 			fmt.Println("⚠️ Umur tidak boleh lebih dari 120!")
// 			continue
// 		}

// 		break
// 	}
// 	fmt.Printf("Hallo %s!\n", nama)
// 	fmt.Printf("Umur kamu %d tahun dan kamu suka bermain %s.", umur, hobi)
// }

// func latihan2() {
// 	reader := bufio.NewReader(os.Stdin)

// 	fmt.Print("Masukan Nilai Anda: ")
// 	nilais, _ := reader.ReadString('\n')
// 	nilais = strings.TrimSpace(nilais)
// 	u, err := strconv.Atoi(nilais)
// 	if err != nil {
// 		fmt.Println("⚠️ Umur harus berupa angka!")
// 	}

// 	var nilai string
// 	if u >= 0 && u <= 100 {
// 		switch {
// 		case u >= 100:
// 			nilai = "A"
// 		case u >= 70:
// 			nilai = "B"
// 		case u >= 50:
// 			nilai = "C"
// 		case u >= 30:
// 			nilai = "D"
// 		case u >= 00:
// 			nilai = "E"
// 		default:
// 			nilai = "Nilai anda kosong"
// 		}
// 	} else {
// 		fmt.Println("Nilai yang anda masukan tidak valid!!")
// 	}
// 	fmt.Println("Nilai Kamu Adalah ", nilai)
// }

// func latihan3() {
// // for dengan counter
// for i := 0; i <= 10; i++ {
// 	fmt.Println("angka kamu adalah ", i)
// }

// // for infinity loop
// for {
// 	fmt.Println("oke")
// 	break //untuk keluar dari infinity loop
// }

// // for dengan continue dan break
// for i := 0; i < 10; i++ {
// 	if i == 5 {
// 		continue
// 	}

// 	if i == 8 {
// 		break
// 	}
// 	fmt.Println(i)
// }

// panggil children

// nama := inputname()
// latihan3child1(nama)

// a := inputangkaA()
// b := inputangkaB()
// hasil := latihan3child2(a, b)
// fmt.Println("hasil dari", a, "ditambah", b, "adalah", hasil)

// 	angkaA := inputangkaA()
// 	angkaB := inputangkaB()
// 	tambah, kali, sisaBagi := latihan3child3(angkaA, angkaB)
// 	fmt.Println("Ditambah :", tambah, "| Dikali", kali, "| Sisa Bagi :", sisaBagi)
// }

// children latihan3
// // child1
// func inputname() string {
// 	var nama string
// 	fmt.Print("Masukkan Nama Anda: ")
// 	fmt.Scanln(&nama)
// 	return nama
// }

// func latihan3child1(nama string) {
// 	fmt.Println("hai " + nama)
// }

// // child2
// func inputangkaA() int {
// 	reader := bufio.NewReader(os.Stdin)

// 	fmt.Print("Masukan angka ke-1 :")
// 	angkas, _ := reader.ReadString('\n')
// 	angkas = strings.TrimSpace(angkas)
// 	a, err := strconv.Atoi(angkas)

// 	if err != nil {
// 		fmt.Println("⚠️ Input harus berupa angka!")
// 		return inputangkaA()
// 	}
// 	return a
// }

// func inputangkaB() int {
// 	reader := bufio.NewReader(os.Stdin)

// 	fmt.Print("Masukan angka ke-2 :")
// 	angkas, _ := reader.ReadString('\n')
// 	angkas = strings.TrimSpace(angkas)
// 	b, err := strconv.Atoi(angkas)

// 	if err != nil {
// 		fmt.Println("⚠️ Input harus berupa angka!")
// 		return inputangkaA()
// 	}
// 	return b
// }

// func latihan3child2(a int, b int) int {
// 	return a + b
// }

// // child3
// func inputangkaA() int {
// 	reader := bufio.NewReader(os.Stdin)

// 	fmt.Println("Masukan angka ke-1 :")
// 	angkaA, _ := reader.ReadString('\n')
// 	angkaA = strings.TrimSpace(angkaA)
// 	a, err := strconv.Atoi(angkaA)

// 	if err != nil {
// 		fmt.Println("⚠️ Input harus berupa angka!")
// 		return inputangkaA()
// 	}

// 	return a
// }

// func inputangkaB() int {
// 	reader := bufio.NewReader(os.Stdin)

// 	fmt.Println("Masukan angka ke-2 :")
// 	angkaB, _ := reader.ReadString('\n')
// 	angkaB = strings.TrimSpace(angkaB)
// 	b, err := strconv.Atoi(angkaB)

// 	if err != nil {
// 		fmt.Println("⚠️ Input harus berupa angka!")
// 		return inputangkaB()
// 	}

// 	return b
// }

// func latihan3child3(a int, b int) (int, int, int) {
// 	return a + b, a * b, a % b
// }

// func latihan4() {
// for i := 0; i <= 10; i++ {
// 	if i%2 == 0 {
// 		fmt.Println("Ini Angka Genap", i)
// 	}

// 	if i%2 == 1 {
// 		fmt.Println("Ini Angka ganjil", i)
// 	}
// }

// reader := bufio.NewReader(os.Stdin)

// fmt.Print("Masukan Nama Anda :")
// nama, _ := reader.ReadString('\n')
// nama = strings.TrimSpace(nama)

// println("hallo", nama)

// a := inputangkaA()
// b := inputangkaB()
// hasil := kali(a, b)
// fmt.Println("hasil dari", a, "x", b, "adalah", hasil)

// luaspersegipanjang
// 	panjang := inputangkaA()
// 	lebar := inputangkaB()
// 	hasil := luaspersegipanjang(panjang, lebar)
// 	fmt.Println("Luas Persegi Panjang Adalah", hasil)
// }

// // children latihan4
// func inputangkaA() int {
// 	reader := bufio.NewReader(os.Stdin)

// 	fmt.Print("Masukan panjang : ")
// 	angkaA, _ := reader.ReadString('\n')
// 	angkaA = strings.TrimSpace(angkaA)
// 	a, err := strconv.Atoi(angkaA)

// 	if err != nil {
// 		fmt.Println("⚠️ Input harus berupa angka!")
// 		return inputangkaA()
// 	}
// 	return a
// }

// func inputangkaB() int {
// 	reader := bufio.NewReader(os.Stdin)

// 	fmt.Print("Masukan lebar : ")
// 	angkaB, _ := reader.ReadString('\n')
// 	angkaB = strings.TrimSpace(angkaB)
// 	b, err := strconv.Atoi(angkaB)

// 	if err != nil {
// 		fmt.Println("⚠️ Input harus berupa angka!")
// 		return inputangkaB()
// 	}
// 	return b
// }

// func kali(a int, b int) int {
// 	return a * b
// }

// func luaspersegipanjang(panjang int, lebar int) int {
// 	return panjang * lebar
// }

// func latihan5() {
// // array
// var buah [2]int
// buah[0] = 10
// fmt.Println(buah[0])

// buah := [3]string{"jeruk", "mangga", "apel"}
// fmt.Println(buah[0])

// for i := 0; i < len(buah); i++ {
// 	fmt.Println("index ke-", i, "value:", buah[i])
// }

// // slice
// var angka = []int{1, 2, 3, 4}
// angka = append(angka, 5)
// fmt.Println(angka[4])

// for i, val := range angka {
// 	fmt.Println("angka ke-", i, "value:", val)
// }

// buah := [5]string{"apel", "jeruk", "mangga", "nanas", "melon"}
// buahapa := buah[0:3]

// for i, Dbuah := range buahapa {
// 	fmt.Println("index:", i, "value:", Dbuah)
// }

// // map
// mahasiswa := map[string]int{
// 	"andi":  90,
// 	"yanto": 80,
// }

// for nama, val := range mahasiswa {
// 	fmt.Printf("name:%s value:%d ", nama, val)
// }

// var nilai [4]int
// nilai[0] = 60
// nilai[1] = 70
// nilai[2] = 80
// nilai[3] = 90

// for i, nilailoop := range nilai {
// 	fmt.Println("index:", i, "value:", nilailoop)
// }

// var hobi = []string{"basket", "voly", "badminton"}
// hobi = append(hobi, "bola", "runing")
// for i, hobies := range hobi {
// 	fmt.Println("index:", i, "values:", hobies)
// }

// gaji := map[string]int{
// 	"putra":  90,
// 	"kucing": 1100,
// 	"GG":     900,
// }

// gaji["daffa"] = 500

// for i, sgaji := range gaji {
// 	fmt.Println("nama", i, "gaji:", sgaji)
// }

// delete(gaji, "putra")
// fmt.Println("..", '\n')

// for i, sgaji := range gaji {
// 	fmt.Println("nama", i, "gaji:", sgaji)
// }

// 	data := make(map[string]int)

// 	for i := 1; i <= 3; i++ {
// 		fmt.Println("\nData ke-", i)

// 		nama := inputName()
// 		nilai := inputNilai()

// 		data[nama] = nilai
// 	}
// 	fmt.Println("\nHasil Data:")
// 	for nama, nilai := range data {
// 		fmt.Printf("Nama: %s, Nilai: %d\n", nama, nilai)
// 	}
// }

// func inputName() string {
// 	reader := bufio.NewReader(os.Stdin)

// 	fmt.Print("Masukan Nama Anda:")
// 	name, _ := reader.ReadString('\n')
// 	name = strings.TrimSpace(name)
// 	if name == "" {
// 		fmt.Println("⚠️ Input tidak boleh kosong!")
// 		return inputName()
// 	}
// 	return name
// }

// func inputNilai() int {
// 	reader := bufio.NewReader(os.Stdin)

// 	fmt.Print("Masukan Nilai Anda:")
// 	nilai, _ := reader.ReadString('\n')
// 	nilai = strings.TrimSpace(nilai)
// 	N, err := strconv.Atoi(nilai)

// 	if err != nil {
// 		fmt.Println("⚠️ Input harus berupa angka!")
// 		return inputNilai()
// 	}
// 	return N
// }

// func latihan6() {
// // struct
// type mahasiswa struct {
// 	Nama  string
// 	umur  int
// 	nilai float64
// }

// var daftar []mahasiswa

// daftar = append(daftar, mahasiswa{"andi", 15, 89.5})
// daftar = append(daftar, mahasiswa{"daffa", 20, 90})

// for _, i := range daftar {
// 	fmt.Println("hai", i.Nama, "umur kamu", i.umur, "dan kamu mendapatkan nilai", i.nilai)
// }

// var mhs mahasiswa
// for _, m := range daftar {
// 	if m.Nama == "andi" {
// 		mhs = m
// 		break
// 	}
// }

// mhs.umur = 40
// fmt.Println(mhs)

// type buku struct {
// 	Judul   string
// 	Penulis string
// 	Harga   int
// }

// b1 := buku{"Penerjemah Bahasa", "Jawa", 50000}
// b2 := buku{"Pemukul Kapak", "yanto", 50000}

// fmt.Println(b1, b2)

// type mahasiswa struct {
// 	Nama    string
// 	Jurusan string
// 	Umur    int
// 	Nilai   float64
// }

// var daftar []mahasiswa
// daftar = append(daftar, mahasiswa{"yanto", "RPL", 50, 80.4})
// daftar = append(daftar, mahasiswa{"yanti", "RPL", 50, 80.4})
// daftar = append(daftar, mahasiswa{"agus", "RPL", 50, 80.43})

// for _, s := range daftar {
// 	fmt.Printf("Nama: %s | Jurusan: %s | Umur: %d | Nilai: %.2f\n", s.Nama, s.Jurusan, s.Umur, s.Nilai)
// }

// 	Data := siswa{"andi", "rpl", 80.2}
// 	fmt.Println("Sebelum", Data.Nilai)
// 	ubahNilai(&Data)
// 	fmt.Println("Sesudah", Data.Nilai)
// }

// type siswa struct {
// 	Nama    string
// 	Jurusan string
// 	Nilai   float64
// }

// func ubahNilai(s *siswa) {
// 	s.Nilai = 100
// }

// func latihanCRUDdatadummy() {
// 	for {
// 		fmt.Println("\n=== MENU ===")
// 		fmt.Println("1.Tambah Mahasiswa")
// 		fmt.Println("2.Tampilkan Mahasiswa")
// 		fmt.Println("3.Update Mahasiswa")
// 		fmt.Println("4.Hapus Mahasiswa")
// 		fmt.Println("5.Keluar")

// 		fmt.Print("Pilih menu: ")
// 		var pilihan int
// 		fmt.Scanln(&pilihan)

// 		switch pilihan {
// 		case 1:
// 			tambahMahasiswa()
// 		case 2:
// 			tampilkanMahasiswa()
// 		case 3:
// 			updateMahasiswa()
// 		case 4:
// 			hapusMahasiswa()
// 		case 5:
// 			fmt.Println("Keluar..")
// 			return
// 		default:
// 			fmt.Println("⚠️ Menu tidak valid.")
// 		}
// 	}
// }

// type mahasiswa struct {
// 	Nama    string
// 	Jurusan string
// 	Nilai   int
// }

// var Daftar []mahasiswa

// // tambah
// func tambahMahasiswa() {
// 	reader := bufio.NewReader(os.Stdin)

// 	fmt.Print("Nama :")
// 	nama, _ := reader.ReadString('\n')
// 	nama = strings.TrimSpace(nama)
// 	if nama == "" {
// 		tambahMahasiswa()
// 	}

// 	fmt.Print("Jurusan :")
// 	Jurusan, _ := reader.ReadString('\n')
// 	Jurusan = strings.TrimSpace(Jurusan)
// 	if Jurusan == "" {
// 		tambahMahasiswa()
// 	}

// 	fmt.Print("Nilai :")
// 	nilai, _ := reader.ReadString('\n')
// 	nilai = strings.TrimSpace(nilai)
// 	n, err := strconv.Atoi(nilai)

// 	if err != nil {
// 		tambahMahasiswa()
// 	}

// 	Daftar = append(Daftar, mahasiswa{nama, Jurusan, n})
// 	fmt.Println("✅ Mahasiswa berhasil ditambahkan.")
// }

// func tampilkanMahasiswa() {
// 	fmt.Println("\n=== Data Mahasiswa ===")
// 	if len(Daftar) == 0 {
// 		fmt.Println("Belum ada data.")
// 		return
// 	}

// 	for i, m := range Daftar {
// 		fmt.Printf("%d.Nama:%s | Jurusan:%s | Nilai: %d\n", i+1, m.Nama, m.Jurusan, m.Nilai)
// 	}
// }

// func updateMahasiswa() {
// 	if len(Daftar) == 0 {
// 		fmt.Println("Belum ada data.")
// 		return
// 	}

// 	fmt.Println("\n=== Data Mahasiswa ===")
// 	if len(Daftar) == 0 {
// 		fmt.Println("Belum ada data.")
// 		return
// 	}

// 	for i, m := range Daftar {
// 		fmt.Printf("%d.Nama:%s | Jurusan:%s | Nilai: %d\n", i+1, m.Nama, m.Jurusan, m.Nilai)
// 	}

// 	reader := bufio.NewReader(os.Stdin)

// 	fmt.Print("Masukkan nama mahasiswa yang ingin diupdate: ")
// 	nama, _ := reader.ReadString('\n')
// 	nama = strings.TrimSpace(nama)

// 	found := false
// 	for i := range Daftar {
// 		if strings.EqualFold(Daftar[i].Nama, nama) {
// 			found = true

// 			fmt.Print("Jurusan Baru:")
// 			jurusan, _ := reader.ReadString('\n')
// 			jurusan = strings.TrimSpace(jurusan)

// 			fmt.Print("Nilai baru: ")
// 			nilaiStr, _ := reader.ReadString('\n')
// 			nilaiStr = strings.TrimSpace(nilaiStr)
// 			nilai, err := strconv.Atoi(nilaiStr)
// 			if err != nil {
// 				fmt.Println("⚠️ Nilai harus berupa angka!")
// 				return
// 			}

// 			Daftar[i].Jurusan = jurusan
// 			Daftar[i].Nilai = nilai

// 			fmt.Println("✅ Data berhasil diupdate.")
// 			break
// 		}
// 	}

// 	if !found {
// 		fmt.Println("❌ Nama mahasiswa tidak ditemukan.")
// 	}
// }

// func hapusMahasiswa() {
// 	fmt.Println("\n=== Data Mahasiswa ===")
// 	if len(Daftar) == 0 {
// 		fmt.Println("Belum ada data.")
// 		return
// 	}

// 	for i, m := range Daftar {
// 		fmt.Printf("%d.Nama:%s | Jurusan:%s | Nilai: %d\n", i+1, m.Nama, m.Jurusan, m.Nilai)
// 	}

// 	reader := bufio.NewReader(os.Stdin)

// 	fmt.Print("Masukkan nama mahasiswa yang ingin dihapus: ")
// 	nama, _ := reader.ReadString('\n')
// 	nama = strings.TrimSpace(nama)

// 	found := false

// 	for i := range Daftar {
// 		if strings.EqualFold(Daftar[i].Nama, nama) {
// 			if i == len(Daftar)-1 {
// 				Daftar = Daftar[:i]
// 			} else {
// 				Daftar = append(Daftar[:i], Daftar[i+1:]...)
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
