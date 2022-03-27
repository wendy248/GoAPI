package models

type MataKuliah struct {
	ID            int8   `json:"id" gorm:"primary_key AUTO_INCREMENT NOT_NULL"`
	KodeMatkul    string `json:"kode matkul" gorm:"primary_key"`
	NamaMatkul    string `json:"nama matkul"`
	JumlahSKS     int16  `json:"jumlah sks"`
	DosenPengampu string `json:"dosen pengampu"`
}
