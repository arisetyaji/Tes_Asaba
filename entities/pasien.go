package entities

type tb_produk struct {
	Id           int
	Nama		 string `validate:"required" label:"Nama Lengkap"`
	NIK          int `validate:"required"`
	
}