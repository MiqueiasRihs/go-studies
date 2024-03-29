package servidor

import (
	"crud/banco"
	"encoding/json"
	"io/ioutil"
	"net/http"
)

type usuario struct {
	ID    uint32 `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

func CriarUsuario(w http.ResponseWriter, r *http.Request) {
	corpoRequisicao, erro := ioutil.ReadAll(r.Body)
	if erro != nil {
		w.Write([]byte("Falha ao ler o corpo da requisição!"))
		return
	}

	var usuario usuario
	if erro = json.Unmarshal(corpoRequisicao, &usuario); erro != nil {
		w.Write([]byte("Erro ao converter o usuário para struct!"))
		return
	}

	db, erro := banco.Conectar()
	if erro != nil {
		w.Write([]byte("Erro ao conectar no banco de dados!"))
		return
	}

	sqlStatement := `INSERT INTO usuarios (nome, email) VALUES ($1, $2)`
	_, erro = db.Exec(sqlStatement, usuario.Nome, usuario.Email)
	if erro != nil {
		panic(erro)
	}

	// statement, erro := db.Prepare("INSERT INTO usuarios (nome, email) VALUES ($1, $2)")
	// if erro != nil {
	// 	w.Write([]byte("Erro ao criar o statement!"))
	// 	log.Fatal(erro)
	// 	return
	// }
	// defer statement.Close()

	// insercao, erro := statement.Exec(usuario.Nome, usuario.Email)
	// if erro != nil {
	// 	w.Write([]byte("Erro ao executar o statement!"))
	// 	return
	// }

	// idInserido, erro := insercao.LastInsertId()
	// if erro != nil {
	// 	w.Write([]byte("Erro ao obter o ID inserido!"))
	// 	return
	// }

	// w.WriteHeader(http.StatusCreated)
	// w.Write([]byte(fmt.Sprintf("Usuário inserido com sucesso! ID: %d", idInserido)))

}
