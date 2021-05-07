package Store

type Conta struct {
	ID        int    `json:"id"`
	Name      string `json:"name"`
	CPF       string `json:"cpf"`
	Secret    string `json:"secret"`
	Balance   int    `json:"balance"`
	CreatedAt string `json:"created_at"`
}

type ArmazenamentoDeContas struct {
	armazenamento map[int]Conta
}
