package models

type Mahasiswa struct {
	ID            int16  `json:"id" gorm:"primary_key AUTO_INCREMENT NOT_NULL"`
	Nama          string `json:"nama"`
	Prodi         string `json:"prodi"`
	Fakultas      string `json:"fakultas"`
	NIM           int64  `json:"nim" gorm:"primary_key NOT_NULL"`
	TahunAngkatan int16  `json:"tahun"`
}
