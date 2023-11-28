package cep

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"net/http"
	"zapys/src/respostas"
)

type Cep struct{
	Cep string `json:"cep"`
	Logradouro string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro string `json:"bairro"`
	Localidade string `json:"localidade"`
	Uf string `json:"uf"`
	Ibde string `json:"ibge"`
	Gia string `json:"gia"`
	Ddd string `json:"ddd"`
	Siafi string `json:"siafi"`
}

func CEP(w http.ResponseWriter, r *http.Request) {
	cep := r.URL.Query().Get("cep")
	if cep == "" {
		respostas.ERROR(w, http.StatusBadRequest, errors.New("Parâmetro cep está vazio"))
		return
	}
	url := fmt.Sprintf("https://viacep.com.br/ws/%s/json/", cep)
	resp, err := http.Get(url)
	if err != nil {
		respostas.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	defer resp.Body.Close()
	if resp.StatusCode != http.StatusOK {
		respostas.ERROR(w, http.StatusInternalServerError, errors.New("Erro ao consultar a API"))
		return
	}
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		respostas.ERROR(w, http.StatusUnsupportedMediaType, err)
		return
	}
	var ceps Cep
	if err := json.Unmarshal(body, &ceps); err != nil {
		respostas.ERROR(w, http.StatusInternalServerError, err)
		return
	}
	respostas.JSON(w, http.StatusOK, ceps)
}
