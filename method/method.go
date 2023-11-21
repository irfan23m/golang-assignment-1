package method

import (
	"assignment-1/function"
	"assignment-1/models"
	"fmt"
	"strings"
)

type ModelPeserta struct {
	models.Peserta
}

func (m ModelPeserta) FindPeserta() {
	var pesertaList = function.DataPeserta()

	for key, v := range pesertaList {
		if strings.ToLower(m.Nama) == strings.ToLower(v.Nama) {
			fmt.Println("ID :", key)
			fmt.Println("nama :", v.Nama)
			fmt.Println("alamat :", v.Alamat)
			fmt.Println("pekerjaan :", v.Pekerjaan)
			fmt.Println("alasan :", v.Alasan)
		}
	}
}
