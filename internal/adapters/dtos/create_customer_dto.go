package dtos

import "errors"

type CustomerDto struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Cpf   string `json:"cpf"`
	Birth string `json:"birth"`
}

func (c *CustomerDto) CheckFilled() error {
	if c.Cpf == "" {
		return errors.New("CPF is required")
	}
	if c.Name == "" {
		return errors.New("name is required")
	}
	if c.Birth == "" {
		return errors.New("birth is required")
	}
	return nil
}
