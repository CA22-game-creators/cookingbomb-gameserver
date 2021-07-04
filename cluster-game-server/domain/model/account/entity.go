package domain

import (
	"github.com/oklog/ulid/v2"
)

type Account struct {
	ID   ID
	Name Name
}

func FromRepository(id, name string) Account {
	return Account{
		ID:   ID(ulid.MustParse(id)),
		Name: Name(name),
	}
}
