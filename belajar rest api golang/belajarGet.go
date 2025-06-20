package main

type Mahasiswa struct {
	NIM   int    `json:"nim"`
	Nama  string `json:"nama"`
	Umur  int    `json:"umur"`
	Nilai int    `json:"nilai"`
}

// // Data dummy
// var dataMahasiswa = []Mahasiswa{
// 	{101, "Andi", 20, 85},
// 	{102, "Budi", 21, 90},
// }

// func hendlermahasiswa(w http.ResponseWriter, r *http.Request) {
// 	w.Header().Set("Content-Type", "application/json")
// 	json.NewEncoder(w).Encode(dataMahasiswa)
// }

// func hello(w http.ResponseWriter, r *http.Request) {
// 	fmt.Fprintln(w, "Selamat datang di API Mahasiswa")
// }

// func status(w http.ResponseWriter, r *http.Request) {
// 	res := map[string]string{"status": "ok"}
// 	w.Header().Set("Content-Type", "appliacation/json")
// 	json.NewEncoder(w).Encode(res)
// }

// func rangedata(w http.ResponseWriter, r *http.Request) {
// 	jumlah := len(dataMahasiswa)
// 	res := map[string]int{"status": jumlah}
// 	w.Header().Set("Content-Type", "appliacation/json")
// 	json.NewEncoder(w).Encode(res)
// }
