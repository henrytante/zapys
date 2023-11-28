package admin

import (
	
	"net/http"
	"zapys/src/db"
	"zapys/src/respostas"
	
)

type Usuario struct{
	ID int `json:"id"`
	Nome string `json:"nome"`
	Token string `json:"token"`
}

func Users(w http.ResponseWriter, r *http.Request)  {
	
	db, err := db.DBToken()
	if err != nil{
		respostas.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()
	rows, err := db.Query("SELECT id,nome,token FROM usuarios")	
	if err != nil{
		respostas.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer rows.Close()
	usuarios := []Usuario{}
	for rows.Next(){
		var usuario Usuario
		if err := rows.Scan(&usuario.ID, &usuario.Nome, &usuario.Token); err != nil{
			respostas.ERROR(w, http.StatusInternalServerError, err)
			return
		}
		usuarios = append(usuarios, usuario)
	}
	if err := rows.Err(); err != nil{
		respostas.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	respostas.JSON(w, http.StatusOK, usuarios)
}