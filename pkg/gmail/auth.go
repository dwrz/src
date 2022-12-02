package gmail

import (
	"fmt"
	"net/smtp"
)

const Address = "smtp.gmail.com:587"

type Auth struct {
	Username string
	Password string
}

func (a *Auth) Start(server *smtp.ServerInfo) (string, []byte, error) {
	return "LOGIN", []byte{}, nil
}

func (a *Auth) Next(fromServer []byte, more bool) ([]byte, error) {
	if more {
		fs := string(fromServer)

		switch fs {
		case "Username:":
			return []byte(a.Username), nil
		case "Password:":
			return []byte(a.Password), nil
		default:
			return nil, fmt.Errorf(
				"unrecognized fromServer: %s", fs,
			)
		}
	}

	return nil, nil
}
