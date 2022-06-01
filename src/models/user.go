package models

import (
	"api/safety"
	"errors"
	"strings"
	"time"

	"github.com/badoux/checkmail"
)

type User struct {
	ID 			uint64		`json:"id,omitempty"`
	Name 		string		`json:"name,omitempty"`
	Nick 		string		`json:"nick,omitempty"`
	Email 		string		`json:"email,omitempty"`
	Password 	string		`json:"password,omitempty"`
	CreatedIn 	time.Time	`json:"createdIn,omitempty"`
}

func (user *User) Prepare(etapa string) error {
	if erro := user.validar(etapa); erro != nil {
		return erro
	}

	if err := user.formatar(etapa); err != nil {
		return err
	}
	return nil
	
}

func (user *User) validar(etapa string) error {
	if user.Name == "" {
		return errors.New("O NOME É OBRIGATORIO E NAO PODE ESTAR EM BRANCO!")
	}

	if user.Nick == "" {
		return errors.New("O NICK É OBRIGATORIO E NAO PODE ESTAR EM BRANCO!")
	}

	if user.Email == "" {
		return errors.New("O EMAIL É OBRIGATORIO E NAO PODE ESTAR EM BRANCO!")
	}

	if err := checkmail.ValidateFormat(user.Email); err != nil {
		return errors.New("O EMAIL INSERIDO E INVALIDO!")
	} 

	if  etapa == "cadastro" && user.Password == "" {
		return errors.New("A SENHA É OBRIGATORIO E NAO PODE ESTAR EM BRANCO!")
	}

	return nil
}

func (user *User) formatar(etapa string) error{
	user.Name = strings.TrimSpace(user.Name)
	user.Nick = strings.TrimSpace(user.Nick)
	user.Email = strings.TrimSpace(user.Email)

	if etapa == "cadastro" {
		passwordComHash, err := safety.Hash(user.Password)
		if err != nil {
			return err
		}

		user.Password = string(passwordComHash)
	}

	return nil
}