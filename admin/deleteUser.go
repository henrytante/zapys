package admin

import (
	
	"errors"
	"fmt"
	"net/http"
	"zapys/src/db"
	"zapys/src/respostas"
)

func DeleteUser(w http.ResponseWriter, r *http.Request) {
	
	id := r.URL.Query().Get("id")
	if id == "" {
		respostas.ERROR(w, http.StatusBadRequest, errors.New("Parâmetro id vazio"))
		return
	}
	dbConn, err := db.DBToken()
	if err != nil {
		respostas.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer dbConn.Close()

	result, err := dbConn.Exec("DELETE FROM usuarios WHERE id = ?", id)
	if err != nil {
		respostas.ERROR(w, http.StatusInternalServerError, err)
		return
	}

	rowsAffected, _ := result.RowsAffected()
	if rowsAffected == 0 {
		respostas.ERROR(w, http.StatusNotFound, errors.New("ID não existente"))
		return
	}

	respostas.JSON(w, http.StatusOK, fmt.Sprintf("Usuário de ID %s foi deletado", id))
}
