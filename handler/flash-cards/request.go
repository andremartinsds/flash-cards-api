package handler

import (
	"fmt"
)

type CreateFlashCardRequest struct {
	Front string `json:"front"`
	Back  string `json:"back"`
	DecId uint   `json:"decId"`
}
type UpdateFlashCardRequest struct {
	Front string `json:"front"`
	Back  string `json:"back"`
	DecId uint   `json:"decId"`
}

func paramRequired(name, typ string) error {

	return fmt.Errorf("param %v (type %s) is required", name, typ)

}

func (d *CreateFlashCardRequest) FlashCardCreateValidate() error {

	if d.Front == "" {
		return paramRequired("name", "string")
	}

	if d.Back == "" {
		return paramRequired("name", "string")
	}

	return nil

}

func (d *UpdateFlashCardRequest) FlashCardUpdateValidate() error {
	if d.Front == "" {
		return paramRequired("name", "string")
	}

	if d.Back == "" {
		return paramRequired("name", "string")
	}

	return nil
}
