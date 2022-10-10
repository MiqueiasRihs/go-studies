package main

import "fmt"

func main() {
	usuario := map[string]string{
		"nome":      "Miquéias",
		"sobrenome": "Rihs",
	}

	fmt.Println(usuario)
	fmt.Println(usuario["nome"])

	usuario2 := map[string]map[string]string{
		"nome": {
			"primeiro": "Miquéias",
			"ultimo":   "Oliveira",
		},
		"curso": {
			"nome":   "ADS",
			"campus": "Campus 1",
		},
	}

	fmt.Println(usuario2)
	delete(usuario2, "curso")
	fmt.Println(usuario2)

	// Adicionar
	usuario2["nascimento"] = map[string]string{
		"dia": "19",
		"mes": "09",
		"ano": "1997",
	}

	fmt.Println(usuario2)

}
