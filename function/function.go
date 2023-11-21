package function

import "assignment-1/models"

func DataPeserta() []models.Peserta {
	pesertaList := []models.Peserta{
		{
			Nama:      "Thomas",
			Alamat:    "Kota A",
			Pekerjaan: "Programmer",
			Alasan:    "Alasan Thomas",
		},
		{
			Nama:      "Irfan",
			Alamat:    "Jakarta",
			Pekerjaan: "programmer",
			Alasan:    "None",
		},
		{
			Nama:      "Tasya",
			Alamat:    "Kota B",
			Pekerjaan: "Programmer",
			Alasan:    "None",
		},
	}

	return pesertaList
}
