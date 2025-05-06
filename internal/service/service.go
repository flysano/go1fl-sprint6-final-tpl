package service

import (
	"errors"
	"strings"

	"go1fl-sprint6-final/pkg/morse"
)

var ErrEmptyOrInvalid = errors.New("Данные некорректны или отсутствуют")

func Converter(data string) (string, error) {
	textOrMorse := func(data string) (string, error) {
		split := strings.Split(data, " ")
		if len(split) == 0 || len(split) == 1 && split[0] == "" {
			return "", ErrEmptyOrInvalid
		}

		for _, s := range split {
			for _, r := range s {
				switch r {
				case '.', '-':
				default:
					return "text", nil
				}
			}
		}
		return "morse", nil
	}

	res, err := textOrMorse(data)
	if err != nil {
		return "", err
	}

	switch res {
	case "morse":
		return morse.ToText(data), nil
	case "text":
		return morse.ToMorse(data), nil
	default:
		return "", ErrEmptyOrInvalid
	}

}
