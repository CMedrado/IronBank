package entries

type Login struct {
	CPF    string `json:"cpf"`
	Secret string `json:"secret"`
}
