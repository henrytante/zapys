package entities

type CPFV2 struct {
	CPF                 string `json:"cpf"`
	PAI                 string `json:"pai"`
	MAE                 string `json:"mae"`
	MunicipioNascimento string `json:"municipioNascimento"`
	Municipio           string `json:"municipio"`
	Logradouro          string `json:"logradouro"`
	Numero              string `json:"numero"`
	Bairro              string `json:"bairro"`
	CEP                 string `json:"cep"`
	RGNumero            string `json:"rgNumero"`
	RGOrgaoEmissor      string `json:"rgOrgaoEmissor"`
	RGUf                string `json:"rgUf"`
	RGDataEmissao       string `json:"rgDataEmissao"`
	CNS                 string `json:"cns"`
	Telefone            string `json:"telefone"`
	TelefoneSecundario  string `json:"telefoneSecundario"`
}

type CPFV1 struct{
	CPF string `json:"cpf"`
	NOME string `json:"nome"`
	SEXO string `json:"sexo"`
	NASCIMENTO string `json:"nascimento"`
}

type NOMEV1 struct{
	CPF string `json:"cpf"`
	NOME string `json:"nome"`
	SEXO string `json:"sexo"`
	NASCIMENTO string `json:"nascimento"`
}