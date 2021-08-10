package authentication

type LoginRequest struct {
	CPF    string `json:"cpf"`
	Secret string `json:"secret"`
}
