package middleware

import (
	"errors"
	"net/http"
	"zapys/src/db"
	"zapys/src/respostas"
)


func ADMMIDDLEWARE(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		admpass := r.URL.Query().Get("admpass")
		if admpass == ""{
			respostas.ERROR(w, http.StatusBadRequest, errors.New("Senha de admin esta em branco"))
			return
		}
		db, err := db.DBADM()
		if err != nil{
			respostas.ERROR(w, http.StatusInternalServerError, err)
			return
		}
		defer db.Close()
		query := "SELECT adminpass FROM admin WHERE adminpass = ?"
		row := db.QueryRow(query, admpass)
		var dbPASS string
		err = row.Scan(&dbPASS)
		if err != nil || dbPASS == ""{
			respostas.ERROR(w, http.StatusUnauthorized, errors.New("Senha invalida"))
			return
		}
		next(w, r)
	}
}