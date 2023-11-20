package handler

import (
	"errors"
	"fmt"
)

type CreateUserRequest struct {
	Name             string `json:"name"`
	Email            string `json:"email"`
	Pass             string `json:"pass"`
	PassConfirmation string `json:"passConfirmation"`
}
type UpdateUserRequest struct {
	Name             string `json:"name"`
	Email            string `json:"email"`
	Pass             string `json:"pass"`
	PassConfirmation string `json:"passConfirmation"`
}

func paramRequired(name, typ string) error {

	return fmt.Errorf("param %v (type %s) is required", name, typ)

}

func validateEmail(email string) error {
	//TODO add email validations
	if email == "" {
		return fmt.Errorf("")
	}
	return nil
}

func isValidPass(pass, passConfirmation string) error {
	if pass == passConfirmation {
		return nil
	}

	return errors.New("The passwords do not match.")
}

func (u *CreateUserRequest) UserCreateValidate() error {
	if err := isValidPass(u.Pass, u.PassConfirmation); err != nil {
		return errors.New("The passwords do not match.")
	}

	if u.Name == "" {
		return paramRequired("name", "string")
	}

	if err := validateEmail(u.Email); err != nil {
		return errors.New("The email needs a value.")
	}

	return nil

}

func (u *UpdateUserRequest) UpdateUserValidate() error {
	if u.Name == "" {
		return paramRequired("name", "string")
	}

	if err := validateEmail(u.Email); err != nil {
		return errors.New("The email needs a value")
	}

	return nil
}
