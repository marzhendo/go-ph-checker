# pH Checker 🧪

Program sederhana untuk mengecek dan mengelola nilai pH produk/larutan dengan interface command-line yang user-friendly.

## ✨ Fitur

- **Cek pH**: Input nama produk dan nilai pH (0-14) dengan validasi otomatis
- **Kategorisasi Otomatis**: Mengelompokkan pH menjadi 5 kategori dengan emoji visual:
  - 🔴 Sangat Asam (0-3.0)
  - 🟠 Asam (3.1-6.4)
  - 🟢 Netral (6.5-7.5)
  - 🔵 Basa (7.6-9.9)
  - 🟣 Sangat Basa (10.0-14.0)
- **Riwayat Lengkap**: Menyimpan semua pengecekan dengan timestamp
- **Persistent Storage**: Data tersimpan dalam file JSON

## 🚀 Cara Penggunaan

```bash
# Clone repository
git clone https://github.com/username/ph-checker.git
cd ph-checker

# Jalankan program
go run main.go
```

## 📝 Menu Program

1. **Cek pH** - Input produk baru dan cek nilai pH
2. **Riwayat pH** - Lihat semua pengecekan sebelumnya
3. **Keluar** - Tutup program

## 🛠️ Teknologi

- **Go** - Bahasa pemrograman utama
- **JSON** - Format penyimpanan data
- **Validasi Input** - Error handling dan retry mechanism

## 📄 Contoh Output

```
--- Hasil Pengecekan pH ---
Nama Produk: Larutan Cuka
Nilai pH: 2.50 🔴
Kategori: Sangat Asam
Waktu Pengecekan: 05-08-2025
-----------------------------
```

