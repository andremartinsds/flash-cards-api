package handler

import (
	"fmt"
)

type Dec struct {
}

type CreateDecRequest struct {
	Name   string `json:"name"`
	UserID uint   `json:"userId"`
}
type UpdateDecRequest struct {
	Name   string `json:"name"`
	UserID uint   `json:"userId"`
}

func paramRequired(name, typ string) error {

	return fmt.Errorf("param %v (type %s) is required", name, typ)

}

func (d *CreateDecRequest) DecCreateValidate() error {

	if d.Name == "" {
		return paramRequired("name", "string")
	}

	if d.UserID == 0 {
		return paramRequired("name", "string")
	}

	return nil

}

func (d *UpdateDecRequest) UpdateDecValidate() error {
	if d.Name == "" {
		return paramRequired("name", "string")
	}

	if d.UserID == 0 {
		return paramRequired("name", "string")
	}

	return nil
}
