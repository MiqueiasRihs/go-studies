package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type cachorro struct {
	Nome  string `json:"nome"`
	Raca  string `json:"raca"`
	Idade uint   `json:"idade"`
}

func main() {
	c := cachorro{"Apolo", "Golden Retriever", 4}

	cachorroEMJSON, erro := json.Marshal(c)
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(cachorroEMJSON)
	fmt.Println(bytes.NewBuffer(cachorroEMJSON))

	c2 := map[string]string{
		"nome": "Laika",
		"raca": "Vira-latinha",
	}

	cachorro2EMJSON, erro := json.Marshal(c2)
	if erro != nil {
		log.Fatal(erro)
	}

	fmt.Println(cachorro2EMJSON)
	fmt.Println(bytes.NewBuffer(cachorro2EMJSON))

}
