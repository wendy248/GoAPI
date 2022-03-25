package models

type Mahasiswa struct {
	ID            int16  `json:"id" gorm:"primary_key"`
	Nama          string `json:"nama"`
	Prodi         string `json:"prodi"`
	Fakultas      string `json:"fakultas"`
	NIM           int16  `json:"nim"`
	TahunAngkatan int16  `json:"tahun"`
}