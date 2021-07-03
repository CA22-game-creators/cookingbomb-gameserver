package domain

import (
	"errors"

	"github.com/oklog/ulid/v2"
)

type Account struct {
	ID   ID
	Name Name
}

func New(id ID, name Name) (Account, error) {
	if (id == ID{}) {
		return Account{}, errors.New("user id is nil")
	}
	if name == "" {
		return Account{}, errors.New("user name is nil")
	}
	return Account{
		ID:   id,
		Name: name,
	}, nil
}

func FromRepository(id, name string) Account {
	return Account{
		ID:   ID(ulid.MustParse(id)),
		Name: Name(name),
	}
}
