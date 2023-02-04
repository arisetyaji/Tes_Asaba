package main

import (
	"net/http"

	"github.com/arisetyaji/Tes_Asaba/controllers/pasiencontroller"
)

func main() {

	http.HandleFunc("/", pasiencontroller.Add)
	http.HandleFunc("/pasien", pasiencontroller.Add)
	http.HandleFunc("/pasien/index", pasiencontroller.Index)
	http.HandleFunc("/pasien/add", pasiencontroller.Add)
	http.HandleFunc("/pasien/delete", pasiencontroller.Delete)

	http.ListenAndServe(":3000", nil)
}