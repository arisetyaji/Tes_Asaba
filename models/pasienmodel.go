package models

import (
	"database/sql"
	"fmt"

	"github.com/arisetyaji/Tes_Asaba/config"
	"github.com/arisetyaji/Tes_Asaba/entities"
)

type PasienModel struct {
	conn *sql.DB
}

func NewPasienModel() *PasienModel {
	conn, err := config.DBConnection()
	if err != nil {
		panic(err)
	}

	return &PasienModel{
		conn: conn,
	}
}

func (p *PasienModel) FindAll() ([]entities.Pasien, error) {

	rows, err := p.conn.Query("select * from pasien")
	if err != nil {
		return []entities.Pasien{}, err
	}
	defer rows.Close()

	var dataPasien []entities.Pasien
	for rows.Next() {
		var pasien entities.Pasien
		rows.Scan(&pasien.Id,
			&pasien.Nama,
			&pasien.Harga)
			

		dataPasien = append(dataPasien, pasien)
	}

	return dataPasien, nil

}

func (p *PasienModel) Create(pasien entities.Pasien) bool {

	result, err := p.conn.Exec("insert into pasien (nama, harga) values(?,?,?)",
		pasien.Nama, pasien.Harga)

	if err != nil {
		fmt.Println(err)
		return false
	}

	lastInsertId, _ := result.LastInsertId()

	return lastInsertId > 0
}

func (p *PasienModel) Find(id int64, pasien *entities.Pasien) error {

	return p.conn.QueryRow("select * from pasien where id = ?", id).Scan(
		&pasien.Id,
		&pasien.Nama,
		&pasien.Harga)
}

func (p *PasienModel) Update(pasien entities.Pasien) error {

	_, err := p.conn.Exec(
		"update pasien set nama = ?, harga = ? ",
		pasien.Nama, pasien.Harga)

	if err != nil {
		return err
	}

	return nil
}

func (p *PasienModel) Delete(id int64) {
	p.conn.Exec("delete from pasien where id = ?", id)
}