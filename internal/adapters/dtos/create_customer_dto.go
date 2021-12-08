package dtos

type CustomerDto struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Cpf   string `json:"cpf"`
	Birth string `json:"birth"`
}
