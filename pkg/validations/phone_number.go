package validations

import (
	"regexp"

	"github.com/severusTI/auth_golang/pkg/ops"
)

func IsValidPhoneNumber(phoneNumber string) error {
	// Verificar se o número de telefone contém apenas dígitos numéricos
	regex := regexp.MustCompile(`^\d+$`)
	if !regex.MatchString(phoneNumber) {
		return ops.NewErro("O número de telefone deve conter apenas dígitos numéricos")
	}

	// Verificar se o número possui comprimento mínimo e máximo
	minLength := 8
	maxLength := 13

	if len(phoneNumber) < minLength || len(phoneNumber) > maxLength {
		return ops.NewErro("O número de telefone deve ter entre 8 e 13 dígitos")
	}

	return nil
}
