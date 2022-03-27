package models

type MataKuliah struct {
	ID            int8   `json:"id" gorm:"primary_key AUTO_INCREMENT NOT_NULL"`
	KodeMatkul    string `json:"kode"`
	NamaMatkul    string `json:"nama"`
	JumlahSKS     int16  `json:"jumlah"`
	DosenPengampu string `json:"dosen"`
}
