package cpfv2

import (
	"database/sql"
	"errors"
	"net/http"
	"zapys/src/db"
	"zapys/src/entities"
	"zapys/src/respostas"
)

func CPFV2(w http.ResponseWriter, r *http.Request) {
	cpf := r.URL.Query().Get("cpf")
	if cpf == "" {
		respostas.ERROR(w, http.StatusBadRequest, errors.New("Parametro cpf esta vazio"))
		return
	}
	db, err := db.ConnectV2()
	if err != nil {
		respostas.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer db.Close()
	var pessoa entities.CPFV2
	if err = db.QueryRow("SELECT COALESCE(cpf, 'não encontrado'),COALESCE(pai, 'não encontrado'),COALESCE(mae, 'não encontrado'),COALESCE(municipioNascimento, 'não encontrado'),COALESCE(municipio, 'não encontrado'),COALESCE(logradouro, 'não encontrado'),COALESCE(numero, 'não encontrado'),COALESCE(bairro, 'não encontrado'),COALESCE(cep, 'não encontrado'),COALESCE(rgNumero, 'não encontrado'),COALESCE(rgOrgaoEmisor, 'não encontrado'),COALESCE(rgUf,'não encontrado'),COALESCE(rgDataEmissao, 'não encontrado'),COALESCE(cns, 'não encontrado'),COALESCE(telefone, 'não encontrado'),COALESCE(telefoneSecundario, 'não encontrado') FROM datasus WHERE cpf = ?", cpf).Scan(&pessoa.CPF,
		&pessoa.PAI,
		&pessoa.MAE,
		&pessoa.MunicipioNascimento,
		&pessoa.Municipio,
		&pessoa.Logradouro,
		&pessoa.Numero,
		&pessoa.Bairro,
		&pessoa.CEP,
		&pessoa.RGNumero,
		&pessoa.RGOrgaoEmissor,
		&pessoa.RGUf,
		&pessoa.RGDataEmissao,
		&pessoa.CNS,
		&pessoa.Telefone,
		&pessoa.TelefoneSecundario); err != nil {
		if err == sql.ErrNoRows {
			respostas.ERROR(w, http.StatusNotFound, errors.New("CPF não encontrado"))
			return
		}
		respostas.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	respostas.JSON(w, http.StatusOK, pessoa)
}
