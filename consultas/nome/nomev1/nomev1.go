package nomev1

import (
	
	"errors"
	"net/http"
	"zapys/src/db"
	"zapys/src/entities"
	"zapys/src/respostas"
)

func NOMEV1(w http.ResponseWriter, r *http.Request)  {
	nome := r.URL.Query().Get("nome")
	if nome == ""{
		respostas.ERROR(w, http.StatusBadRequest, errors.New("Parametro nome esta vazio"))
		return
	}
	db, err := db.ConnectV1()
	if err != nil{
		respostas.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()
	rows, err := db.Query("SELECT cpf, nome, sexo, nascimento FROM pessoas WHERE nome = ?", nome)
	if err != nil {
		respostas.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer rows.Close()
	var pessoas []entities.NOMEV1
	for rows.Next(){
		var pessoa entities.NOMEV1
		if err := rows.Scan(&pessoa.CPF, &pessoa.NOME, &pessoa.SEXO, &pessoa.NASCIMENTO); err != nil {
			respostas.ERROR(w, http.StatusInternalServerError, err)
			return
		}
		pessoas = append(pessoas, pessoa)
	}
	if len(pessoas) <= 0{
		respostas.ERROR(w, http.StatusNotFound, errors.New("Nome não encontrado"))
		return
	}
	respostas.JSON(w, http.StatusOK, pessoas)
}