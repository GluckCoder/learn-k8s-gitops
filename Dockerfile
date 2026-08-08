# 1. Ambil sistem operasi mini yang sudah ada Golang-nya dari internet
FROM golang:1.26-alpine

# 2. Buat folder kerja baru di dalam sistem mini tersebut bernama /app
WORKDIR /app

# 3. Salin file main.go dari laptop kita ke dalam folder /app tersebut
COPY main.go .

# 4. Beritahu sistem bahwa aplikasi ini akan menggunakan pintu (port) 8080
EXPOSE 8080

# 5. Instruksi perintah yang dijalankan saat aplikasi dinyalakan
CMD ["go", "run", "main.go"]