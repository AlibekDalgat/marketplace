package validation

import (
	"errors"
	"regexp"
)

func ValidLogin(login string) error {
	if len(login) < 3 || len(login) > 20 {
		return errors.New("Длина логина должен быть длинее 3 и короче 20")
	}

	if match, _ := regexp.MatchString(`^[a-zA-Z0-9-_]+$`, login); !match {
		return errors.New("Логин должен состоять из латинских букв, цифр, дефисов или нижних подчёркиваний")
	}

	return nil
}

func ValidPassword(password string) error {
	if len(password) < 8 && len(password) > 255 {
		return errors.New("Длина пароля должна быть больше 8 и меньше 255")
	}
	if match, _ := regexp.MatchString(`[0-9]`, password); !match {
		return errors.New("Пароль должен содержать хотя бы одну цифру")
	}
	if match, _ := regexp.MatchString(`[a-z]`, password); !match {
		return errors.New("Пароль должен содержать хотя бы одну латинскую букву нижнего регистра")
	}
	if match, _ := regexp.MatchString(`[A-Z]`, password); !match {
		return errors.New("Пароль должен содержать хотя бы одну латинскую букву верхнего регистра")
	}
	if match, _ := regexp.MatchString(`[!@#\$%^&*]`, password); !match {
		return errors.New("Пароль должен содержать хотя бы один специальный символ !, @, #, $, %, &, *")
	}
	return nil
}
