package entities

type tb_produk struct {
	Id           int
	Nama		 string `validate:"required" label:"Nama Lengkap"`
	Ongkir       int `validate:"required" label:"Ongkir"`
	
}