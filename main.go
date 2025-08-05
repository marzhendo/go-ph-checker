package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"
)

const (
	minpH       = 0.0
	maxpH       = 14.0
	maxAttempts = 3
)

type ProdukPH struct {
	Nama     string  `json:"nama"`
	PH       float64 `json:"ph"`
	Kategori string  `json:"kategori"`
	Emoji    string  `json:"emoji"`
	Waktu    string  `json:"waktu"`
}

var riwayatProduk []ProdukPH

func header() {
	fmt.Println("||===========================||")
	fmt.Println("||                           ||")
	fmt.Println("||      Program cek pH       ||")
	fmt.Println("||                           ||")
	fmt.Println("||===========================||")
}

func tampilanMenu() {
	fmt.Println("Menu:")
	fmt.Println("1. Cek pH")
	fmt.Println("2. Riwayat pH")
	fmt.Println("3. Keluar")
	fmt.Print("Pilih menu (1-3): ")
}

func getKategoriPH(pH float64) string {
	switch true {
	case pH >= 0 && pH <= 3.0:
		return "Sangat Asam"
	case pH > 3.0 && pH < 6.5:
		return "Asam"
	case pH >= 6.5 && pH <= 7.5:
		return "Netral"
	case pH > 7.5 && pH < 10.0:
		return "Basa"
	case pH >= 10.0 && pH <= 14.0:
		return "Sangat Basa"
	default:
		return "Invalid. Harus antara 0 dan 14."
	}
}

func getEmojiKategori(kategori string) string {
	switch {
	case kategori == "Sangat Asam":
		return "🔴"
	case kategori == "Asam":
		return "🟠"
	case kategori == "Netral":
		return "🟢"
	case kategori == "Basa":
		return "🔵"
	case kategori == "Sangat Basa":
		return "🟣"
	default:
		return "❓"

	}
}

func inputNamaProduk() string {
	fmt.Print("Masukkan nama produk/larutan: ")

	reader := bufio.NewReader(os.Stdin)
	for {
		var nama string
		var err error
		nama, err = reader.ReadString('\n')
		if err != nil {
			log.Fatal("Gagal membaca input:", err)
		}

		nama = strings.TrimSpace(nama)
		if nama == "" {
			fmt.Print("Nama tidak boleh kosong. Silakan coba lagi: ")
			continue
		}
		return nama
	}

}

func inputPH() (float64, error) {
	var pH float64
	for i := 0; i < maxAttempts; i++ {
		fmt.Print("Masukkan nilai pH (0-14): ")
		_, err := fmt.Scan(&pH)
		if err != nil {
			fmt.Println("Input tidak valid. Pastikan Anda memasukkan angka.")
			bufio.NewReader(os.Stdin).ReadString('\n')
			continue
		}
		if pH < minpH || pH > maxpH {
			fmt.Printf("Nilai pH harus antara %.1f dan %.1f. Silakan coba lagi.\n", minpH, maxpH)
			continue
		}
		return pH, nil
	}
	return 0, fmt.Errorf("gagal memasukkan nilai pH setelah %d percobaan", maxAttempts)
}

func jalankanCekPH() {
	fmt.Println("\n--- Jalankan Cek pH ---")
	nama := inputNamaProduk()
	pH, err := inputPH()

	if err != nil {
		fmt.Println("Error:", err)
		fmt.Println("Kembali ke menu utama.")
		return
	}
	kategori := getKategoriPH(pH)
	emoji := getEmojiKategori(kategori)
	waktuCek := time.Now().Format("02-01-2006")

	produkBaru := ProdukPH{
		Nama:     nama,
		PH:       pH,
		Kategori: kategori,
		Emoji:    emoji,
		Waktu:    waktuCek,
	}

	riwayatProduk = append(riwayatProduk, produkBaru)
	if err := simpanRiwayatKeFile(); err != nil {
		fmt.Println("\nWarning: Gagal menyimpan riwayat ke file.", err)
	}
	fmt.Println("\n--- Hasil Pengecekan pH ---")
	fmt.Printf("Nama Produk: %s\n", nama)
	fmt.Printf("Nilai pH: %.2f %s\n", pH, emoji)
	fmt.Printf("Kategori: %s\n", kategori)
	fmt.Printf("Waktu Pengecekan: %s\n", waktuCek)
	fmt.Println("-----------------------------")
}

func tampilkanRiwayat() {
	if len(riwayatProduk) == 0 {
		fmt.Println("Riwayat pH kosong. Silakan lakukan pengecekan pH terlebih dahulu.")
		return
	}

	for i, produk := range riwayatProduk {
		emoji := getEmojiKategori(produk.Kategori)
		fmt.Printf("Riwayat %d:\n", i+1)
		fmt.Printf("Nama Produk: %s\n", produk.Nama)
		fmt.Printf("Nilai pH: %.2f %s\n", produk.PH, emoji)
		fmt.Printf("Kategori: %s\n", produk.Kategori)
		fmt.Printf("Waktu Pengecekan: %s\n", produk.Waktu)
		fmt.Println("-----------------------------")
	}
}

func simpanRiwayatKeFile() error {
	data, err := json.MarshalIndent(riwayatProduk, "", "  ")
	if err != nil {
		return fmt.Errorf("gagal mengubah data ke JSON: %w", err)
	}
	err = os.WriteFile("riwayat_ph.json", data, 0644)
	if err != nil {
		return fmt.Errorf("gagal menyimpan riwayat ke file: %w", err)
	}
	return nil
}

func muatRiwayatDariFile() error {
	data, err := os.ReadFile("riwayat_ph.json")
	if err != nil {
		if os.IsNotExist(err) {
			return nil // File tidak ada, tidak perlu error
		}
		return fmt.Errorf("gagal membaca file riwayat: %w", err)
	}
	if len(data) == 0 {
		return nil
	}
	err = json.Unmarshal(data, &riwayatProduk)
	if err != nil {
		return fmt.Errorf("gagal mengubah JSON ke data: %w", err)
	}
	return nil

}

func main() {
	if err := muatRiwayatDariFile(); err != nil {
		log.Fatalf("Error saat memuat riwayat dari file: %v", err)
	}

	reader := bufio.NewReader(os.Stdin)
	for {
		header()
		tampilanMenu()

		input, _ := reader.ReadString('\n')
		input = strings.TrimSpace(input)

		pilihan, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Input tidak valid. Harap masukakan angka 1-3.")
			continue
		}

		switch pilihan {
		case 1:
			jalankanCekPH()
		case 2:
			tampilkanRiwayat()
		case 3:
			fmt.Println("Terima kasih telah menggunakan program ini. Sampai jumpa!")
			os.Exit(0)
		default:
			fmt.Println("Pilihan tidak valid. Silakan pilih menu 1-3.")
		}

		fmt.Println("\nTekan Enter untuk melanjutkan...")
		reader.ReadString('\n')
	}
}
