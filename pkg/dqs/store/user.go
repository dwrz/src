package store

import (
	"encoding/json"
	"fmt"
	"os"

	"code.dwrz.net/src/pkg/dqs/user"
)

const userFile = "user.json"

func (s *Store) GetUser() (*user.User, error) {
	name := fmt.Sprintf("%s/%s", s.Dir, userFile)

	data, err := os.ReadFile(name)
	if err != nil {
		return nil, err
	}

	var u = &user.User{}
	if err := json.Unmarshal(data, u); err != nil {
		return nil, err
	}

	return u, nil
}

func (s *Store) UpdateUser(u *user.User) error {
	data, err := json.Marshal(u)
	if err != nil {
		return err
	}

	name := fmt.Sprintf("%s/%s", s.Dir, userFile)

	if err := os.WriteFile(name, data, permissions); err != nil {
		return err
	}

	return nil
}
