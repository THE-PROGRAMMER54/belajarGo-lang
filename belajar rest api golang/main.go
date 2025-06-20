package main

import (
	"fmt"
	"net/http"
	"restApi/crud"
)

func main() {
	// // belajarGet
	// http.HandleFunc("/api/", hello)
	// http.HandleFunc("/api/mahasiswa", hendlermahasiswa)
	// http.HandleFunc("/api/ping", status)
	// http.HandleFunc("/api/jumlah", rangedata)

	// belajarPost
	http.HandleFunc("/api/pmahasiswa", crud.Mahasiswahandler)

	fmt.Println("Server berjalan di http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
