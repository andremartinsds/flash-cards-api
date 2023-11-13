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

func isValidePass(pass, passConfirmation string) error {
	if pass == passConfirmation {
		print("entrou aqui")
		return nil
	}

	return errors.New("the password does not equal")
}

func (u *CreateUserRequest) UserCreateValidate() error {
	if err := isValidePass(u.Pass, u.PassConfirmation); err != nil {
		return errors.New("the password does not equal")
	}

	if u.Name == "" {
		return paramRequired("name", "string")
	}

	if err := validateEmail(u.Email); err != nil {
		return errors.New("The email need a value")
	}

	return nil

}
