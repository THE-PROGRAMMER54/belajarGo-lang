package crud

import (
	"encoding/json"
	"net/http"
	"strconv"
	"strings"
)

type mahasiswa struct {
	NIM   int    `json:"nim"`
	Nama  string `json:"nama"`
	Umur  int    `json:"umur"`
	Nilai int    `json:"nilai"`
}

var dataMahasiswa = []mahasiswa{
	{101, "Andi", 20, 85},
	{102, "Budi", 21, 90},
}

func Mahasiswahandler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
	case http.MethodGet:
		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(dataMahasiswa)

	case http.MethodPost:
		contentType := r.Header.Get("Content-Type")
		var mhs mahasiswa

		if strings.HasPrefix(contentType, "application/json") {
			if err := json.NewDecoder(r.Body).Decode(&mhs); err != nil {
				http.Error(w, "❌ JSON tidak valid", http.StatusBadRequest)
				return
			}

		} else if strings.HasPrefix(contentType, "application/x-www-form-urlencoded") {
			if err := r.ParseForm(); err != nil {
				http.Error(w, "❌ Gagal parsing form", http.StatusBadRequest)
				return
			}

			nim, err1 := strconv.Atoi(r.FormValue("nim"))
			umur, err2 := strconv.Atoi(r.FormValue("umur"))
			nilai, err3 := strconv.Atoi(r.FormValue("nilai"))
			if err1 != nil || err2 != nil || err3 != nil {
				http.Error(w, "❌ NIM, Umur, dan Nilai harus angka", http.StatusBadRequest)
				return
			}

			mhs = mahasiswa{
				NIM:   nim,
				Nama:  r.FormValue("nama"),
				Umur:  umur,
				Nilai: nilai,
			}

		} else if strings.HasPrefix(contentType, "multipart/form-data") {
			if err := r.ParseMultipartForm(10 << 20); err != nil {
				http.Error(w, "❌ Gagal parsing form-data", http.StatusBadRequest)
				return
			}

			nim, err1 := strconv.Atoi(r.FormValue("nim"))
			umur, err2 := strconv.Atoi(r.FormValue("umur"))
			nilai, err3 := strconv.Atoi(r.FormValue("nilai"))
			if err1 != nil || err2 != nil || err3 != nil {
				http.Error(w, "❌ NIM, Umur, dan Nilai harus angka", http.StatusBadRequest)
				return
			}

			mhs = mahasiswa{
				NIM:   nim,
				Nama:  r.FormValue("nama"),
				Umur:  umur,
				Nilai: nilai,
			}

		} else {
			http.Error(w, "❌ Tipe konten tidak didukung", http.StatusUnsupportedMediaType)
			return
		}

		dataMahasiswa = append(dataMahasiswa, mhs)

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"message": "✅ Mahasiswa ditambahkan (dari form-data atau JSON)",
		})

	case http.MethodPut:
		contentType := r.Header.Get("Content-Type")
		var mhs mahasiswa

		if strings.HasPrefix(contentType, "application/json") {
			if err := json.NewDecoder(r.Body).Decode(&mhs); err != nil {
				http.Error(w, "❌ JSON tidak valid", http.StatusBadRequest)
				return
			}

		} else if strings.HasPrefix(contentType, "application/x-www-form-urlencoded") {
			if err := r.ParseForm(); err != nil {
				http.Error(w, "❌ Gagal parsing form", http.StatusBadRequest)
				return
			}

			nim, err1 := strconv.Atoi(r.FormValue("nim"))
			umur, err2 := strconv.Atoi(r.FormValue("umur"))
			nilai, err3 := strconv.Atoi(r.FormValue("nilai"))
			if err1 != nil || err2 != nil || err3 != nil {
				http.Error(w, "❌ NIM, Umur, dan Nilai harus angka", http.StatusBadRequest)
				return
			}

			mhs = mahasiswa{
				NIM:   nim,
				Nama:  r.FormValue("nama"),
				Umur:  umur,
				Nilai: nilai,
			}

		} else if strings.HasPrefix(contentType, "multipart/form-data") {
			if err := r.ParseMultipartForm(10 << 20); err != nil {
				http.Error(w, "❌ Gagal parsing form-data", http.StatusBadRequest)
				return
			}

			nim, err1 := strconv.Atoi(r.FormValue("nim"))
			umur, err2 := strconv.Atoi(r.FormValue("umur"))
			nilai, err3 := strconv.Atoi(r.FormValue("nilai"))
			if err1 != nil || err2 != nil || err3 != nil {
				http.Error(w, "❌ NIM, Umur, dan Nilai harus angka", http.StatusBadRequest)
				return
			}

			mhs = mahasiswa{
				NIM:   nim,
				Nama:  r.FormValue("nama"),
				Umur:  umur,
				Nilai: nilai,
			}

		} else {
			http.Error(w, "❌ Tipe konten tidak didukung", http.StatusUnsupportedMediaType)
			return
		}

		found := false
		for i, m := range dataMahasiswa {
			if m.NIM == mhs.NIM {
				dataMahasiswa[i] = mhs
				found = true
				break
			}
		}

		if !found {
			http.Error(w, "❌ Mahasiswa tidak ditemukan", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"message": "✅ Data mahasiswa berhasil diupdate",
		})

	case http.MethodDelete:
		contentType := r.Header.Get("Content-Type")

		var nim int
		var err error

		if strings.HasPrefix(contentType, "application/json") {
			var req struct {
				NIM int `json:"nim"`
			}
			if err = json.NewDecoder(r.Body).Decode(&req); err != nil {
				http.Error(w, "❌ JSON tidak valid", http.StatusBadRequest)
				return
			}
			nim = req.NIM

		} else if strings.HasPrefix(contentType, "application/x-www-form-urlencoded") {
			if err = r.ParseForm(); err != nil {
				http.Error(w, "❌ Gagal parsing form", http.StatusBadRequest)
				return
			}
			nim, err = strconv.Atoi(r.FormValue("nim"))

		} else if strings.HasPrefix(contentType, "multipart/form-data") {
			if err = r.ParseMultipartForm(10 << 20); err != nil {
				http.Error(w, "❌ Gagal parsing form-data", http.StatusBadRequest)
				return
			}
			nim, err = strconv.Atoi(r.FormValue("nim"))

		} else {
			http.Error(w, "❌ Tipe konten tidak didukung", http.StatusUnsupportedMediaType)
			return
		}

		if err != nil {
			http.Error(w, "❌ NIM tidak valid (harus angka)", http.StatusBadRequest)
			return
		}

		found := false
		for i, mhs := range dataMahasiswa {
			if mhs.NIM == nim {
				dataMahasiswa = append(dataMahasiswa[:i], dataMahasiswa[i+1:]...)
				found = true
				break
			}
		}

		if !found {
			http.Error(w, "❌ Mahasiswa tidak ditemukan", http.StatusNotFound)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		json.NewEncoder(w).Encode(map[string]string{
			"message": "✅ Mahasiswa berhasil dihapus",
		})

	default:
		http.Error(w, "❌ Method tidak dikenali", http.StatusMethodNotAllowed)
	}
}
