package validations

import (
	"errors"
	"regexp"
)

func IsValidEmail(email string) error {
	// Utilize uma expressão regular para validar o formato do e-mail
	// Você pode usar uma expressão regular mais complexa para validar completamente o formato do e-mail
	// Aqui está um exemplo simples apenas para ilustração
	emailRegex := regexp.MustCompile(`^[a-zA-Z0-9._%+-]+@[a-zA-Z0-9.-]+\.[a-zA-Z]{2,}$`)
	if !emailRegex.MatchString(email) {
		return errors.New("invalid email")
	}

	return nil
}
