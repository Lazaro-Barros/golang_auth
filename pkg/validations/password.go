package validations

import (
	"regexp"
	"strings"

	"github.com/severusTI/auth_golang/pkg/ops"
)

func IsValidPassword(password string) error {
	// Verificar o comprimento mínimo
	if len(password) < 8 {
		return ops.NewErro("A senha deve ter no mínimo 8 caracteres")
	}

	// Verificar se a senha contém pelo menos um caractere especial
	hasSpecialChar := regexp.MustCompile(`[!@#$%^&*(),.?":{}|<>]`).MatchString(password)
	if !hasSpecialChar {
		return ops.NewErro("A senha deve conter pelo menos um caractere especial (!@#$%^&*(),.?\":{}|<>)")
	}

	// Verificar se a senha contém pelo menos uma letra maiúscula e uma minúscula
	hasUppercase := false
	hasLowercase := false
	for _, char := range password {
		if strings.ToUpper(string(char)) == string(char) {
			hasUppercase = true
		} else if strings.ToLower(string(char)) == string(char) {
			hasLowercase = true
		}
	}
	if !hasUppercase || !hasLowercase {
		return ops.NewErro("A senha deve conter pelo menos uma letra maiúscula e uma minúscula")
	}

	// Verificar se a senha contém pelo menos um dígito numérico
	hasDigit := regexp.MustCompile(`\d`).MatchString(password)
	if !hasDigit {
		return ops.NewErro("A senha deve conter pelo menos um dígito numérico")
	}

	// A senha atende a todos os critérios
	return nil
}
