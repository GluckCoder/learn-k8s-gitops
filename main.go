package main

import (
	"fmt"
	"net/http"
)

func main() {
	// Membuat endpoint utama "/" yang akan menampilkan teks saat diakses
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Halo! Aplikasi Golang pertama saya berhasil berjalan! 🚀")
	})

	// Menampilkan status di terminal bahwa aplikasi aktif
	fmt.Println("Server sedang berjalan di port 8080...")
	
	// Menyalakan server agar mendengarkan port 8080
	http.ListenAndServe(":8080", nil)
}