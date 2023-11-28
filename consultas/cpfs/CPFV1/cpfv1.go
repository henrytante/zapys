package cpfv1

import (
	"database/sql"
	"errors"
	"net/http"
	"zapys/src/db"
	"zapys/src/entities"
	"zapys/src/respostas"
)

func CPFV1(w http.ResponseWriter, r *http.Request)  {
	cpf := r.URL.Query().Get("cpf")
	if cpf == ""{
		respostas.ERROR(w, http.StatusBadRequest, errors.New("Parametro cpf esta vazio"))
		return
	}
	db, err := db.ConnectV1()
	if err != nil{
		respostas.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()
	var pessoa entities.CPFV1
	if err = db.QueryRow("SELECT cpf,nome,sexo,nascimento from pessoas where cpf = ?", cpf).Scan(&pessoa.CPF, &pessoa.NOME, &pessoa.SEXO, &pessoa.NASCIMENTO); err != nil{
		if err == sql.ErrNoRows{
			respostas.ERROR(w, http.StatusNotFound, errors.New("CPF não encontrado"))
			return
		}
		respostas.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	respostas.JSON(w, http.StatusOK, pessoa)
}