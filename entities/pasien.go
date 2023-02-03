package entities

type Pasien struct {
	Id           int64
	Nama		 string `validate:"required" label:"Nama"`
	Harga        string `validate:"required"`
	
}