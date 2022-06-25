package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	personas := []string{
		"Duna",
		"Eki",
		"Lluvia",
		"Indar",
		"Aran",
		"Markel",
		"Naike",
		"David",
		"Amaia",
		"Mertxe",
		"Inko",
		"Mikel",
		"Amama",
		"Aitata",
	}
	min := 0
	max := 13
	rand.Seed(time.Now().UnixNano())
	azarPersona := rand.Intn(max-min+1) + min

	cosas := []string{
		"haca caca",
		"hace pis",
		"se tira un pedo",
		"vomita",
		"se cae",
		"está leyendo",
		"está durmiendo",
		"hace gimnasia",
		"se pelea",
		"echa un pulso",
		"come gusanitos",
		"se rie",
		"se enfada",
		"llora",
	}
	rand.Seed(time.Now().UnixNano())
	azarCosa := rand.Intn(max-min+1) + min

	sitios := []string{
		"en el balcón",
		"en la cocina",
		"en el baño",
		"en la sala",
		"en el cuarto",
		"en el restaurante",
		"en el parque",
		"en la ikastola",
		"en el monte",
		"en la bicicleta",
		"en el comedor",
		"en el sofá",
		"en la cama",
		"en el suelo",
	}
	rand.Seed(time.Now().UnixNano())
	azarSitio := rand.Intn(max-min+1) + min

	fmt.Println("===========================================================")
	fmt.Println(
		fmt.Sprintf(
			"%s %s %s",
			personas[azarPersona],
			cosas[azarCosa],
			sitios[azarSitio],
		),
	)
	fmt.Println("===========================================================")
}
