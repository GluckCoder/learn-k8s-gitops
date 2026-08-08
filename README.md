# instal kind (simulasi kubernetes)

### 1. Unduh binary kind versi terbaru untuk Linux:
```curl -Lo ./kind https://kind.sigs.k8s.io/dl/latest/kind-linux-amd64```

### 2. Berikan izin hak akses agar file bisa dieksekusi
```chmod +x ./kind```

### 3. Pindahkan ke folder sistem agar bisa dipanggil dari mana saja
```sudo mv ./kind /usr/local/bin/kind```

### 4. Verifikasi bahwa instalasi berhasil dengan mengecek versinya:
```kind version```

### 5. Buat klaster Kubernetes lokal pertama Anda menggunakan perintah:
```
kind create cluster
(Catatan: Jika saat membuat klaster Anda tidak memberi nama khusus, maka nama bawaannya adalah kind). Proses ini akan memakan waktu beberapa detik untuk menyalin file.
```


# Instal Kubectl (Remote Kontrol K8s)

### 1. Unduh binary kubectl resmi untuk arsitektur Linux x86_64:
```curl -LO "https://dl.k8s.io/release/$(curl -L -s https://dl.k8s.io/release/stable.txt)/bin/linux/amd64/kubectl"```

### 2. Berikan izin eksekusi pada file tersebut:
```chmod +x kubectl```

### 3. Pindahkan file ke direktori /usr/local/bin:
```sudo mv ./kubectl /usr/local/bin/kubectl```


## tahap untuk dockerfile
### 1. setelah buat main.go dan go mod init belajar, selanjutnya build image:
```docker build -t golang-pemula:v1 .```
(Catatan penting: Jangan lupakan tanda titik (.) di akhir perintah tersebut).

### 2. testing image : 
```docker run -d -p 8080:8080 --name uji-coba-golang golang-pemula:v1```

### 3. docker ps untuk liat :
```docker ps```

### 4. untuk bersihkan container : 
```docker rm -f uji-coba-golang```

## tahap untuk daaftarin image golang ke dadlam klaster kube
### 1. command : 
```kind load docker-image golang-pemula:v1 --name kind```

### 2. Buat File Instruksi Kubernetes (deployment.yaml)
Di dalam Kubernetes, kita tidak menggunakan perintah `docker run`. Kita menyuruh Kubernetes menggunakan file teks instruksi berformat .yaml.

### 3. setelah buat deployment.yaml kita terbangkan aplikasi ke kube : 
```kubectl apply -f deployment.yaml```

catatan : Di dalam klaster simulasi Kind, aplikasi kita berada di dalam jaringan isolasi Kubernetes yang belum terhubung langsung ke Windows. Kita harus membukakan jalan pintas menggunakan fitur bernama port-forward.

### 4. menghubungkaan pintu kube ke pc kita (port forward)
```kubectl port-forward svc/golang-service 8080:8080```
untuk mematikannya lagi tinggal tekan Ctrl+C

### 5. klau mau udahan, hapus cluster : 
```kind delete cluster --name kind```


1. kind create cluster --name kind
2. kubectl get nodes
(bungkus kode main.go menjadi image)
3. docker build -t golang-pemula:v1 .
4. kind load docker-image golang-pemula:v1 --name kind
