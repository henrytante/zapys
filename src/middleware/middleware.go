package middleware

import (
	"errors"
	"net/http"
	"zapys/src/db"
	"zapys/src/respostas"
)


func TokenMiddleware(next http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		token := r.URL.Query().Get("token")
		db, err := db.DBToken()
		if err != nil{
			respostas.ERROR(w, http.StatusInternalServerError, err)
			return
		}
		defer db.Close()
		query := "SELECT token FROM usuarios WHERE token = ?"
		row := db.QueryRow(query, token)
		var dbToken string
		err = row.Scan(&dbToken)
		if err != nil || dbToken == ""{
			respostas.ERROR(w, http.StatusUnauthorized, errors.New("Token invalido"))
			return
		}
		next(w, r)
	}
}