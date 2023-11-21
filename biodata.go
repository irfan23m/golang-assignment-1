package main

import (
	"assignment-1/method"
	"assignment-1/models"
	"fmt"
	"os"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage : main <name>")
		return
	}

	args := os.Args[1]
	peserta := method.ModelPeserta{
		Peserta: models.Peserta{
			Nama: args,
		},
	}

	peserta.FindPeserta()
}
