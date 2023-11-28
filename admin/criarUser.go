package admin

import (
	"errors"
	"fmt"
	"math/rand"
	"net/http"
	"time"
	"zapys/src/db"
	"zapys/src/respostas"
)

func gerarToken(tamanho int) string {
	var combinação []rune
	rand.Seed(time.Now().UnixNano())

	for i := 0; i < tamanho; i++ {
		var c rune
		switch rand.Intn(3) {
		case 0:
			c = rune(rand.Intn(26) + 'a') // Letras minúsculas
		case 1:
			c = rune(rand.Intn(26) + 'A') // Letras maiúsculas
		case 2:
			c = rune(rand.Intn(10) + '0') // Números
		}
		combinação = append(combinação, c)
	}

	return string(combinação)
}
func CriarUser(w http.ResponseWriter, r *http.Request) {
	username := r.URL.Query().Get("username")
	if username == "" {
		respostas.ERROR(w, http.StatusBadRequest, errors.New("Username esta em branco"))
		return
	}
	token := gerarToken(8)
	db, err := db.DBToken()
	if err != nil {
		respostas.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()
	_, criar := db.Exec("insert into usuarios (nome, token) values (?,?)", username, token)
	if criar != nil {
		respostas.ERROR(w, http.StatusInternalServerError, errors.New("Erro ao criar usuario ou usuario ja existente"))
		return
	}
	respostas.JSON(w, http.StatusCreated, fmt.Sprintf("usuario '%s' criado com sucesso. Token: %s", username, token))
}
