package main
import "fmt"

func main() {
	const (
		kuis float64 = 0.15
		tugas float64 = 0.15
		praktikum float64 = 0.2
		UTS float64 = 0.25
		UAS float64 = 0.25
	)
	type mahasiswa struct {
		NIM string
		nama string
		nilaiKuis float64
		nilaiTugas float64
		nilaiPraktikum float64
		nilaiUTS float64
		nilaiUAS float64
		nilaiAkhir float64
	}
	var DataMahasiswa mahasiswa
	fmt.Scan(&DataMahasiswa.nama, &DataMahasiswa.NIM, &DataMahasiswa.nilaiKuis, &DataMahasiswa.nilaiTugas, &DataMahasiswa.nilaiPraktikum, &DataMahasiswa.nilaiUTS, &DataMahasiswa.nilaiUAS)
	DataMahasiswa.nilaiAkhir = (kuis * DataMahasiswa.nilaiKuis) + (tugas * DataMahasiswa.nilaiTugas) + (praktikum * DataMahasiswa.nilaiPraktikum) + (UTS * DataMahasiswa.nilaiUTS) + (UAS * DataMahasiswa.nilaiUAS)
	fmt.Println("--- Data Mahasiswa ---")
	fmt.Println("Nama:", DataMahasiswa.nama)
	fmt.Println("NIM:", DataMahasiswa.NIM)
	fmt.Println("Nilai Kuis:", DataMahasiswa.nilaiKuis)
	fmt.Println("Nilai Tugas:", DataMahasiswa.nilaiTugas)
	fmt.Println("Nilai Praktikum:", DataMahasiswa.nilaiPraktikum)
	fmt.Println("Nilai UTS:", DataMahasiswa.nilaiUTS)
	fmt.Println("Nilai UAS:", DataMahasiswa.nilaiUAS)
	fmt.Println("Nilai Akhir:", DataMahasiswa.nilaiAkhir)
}