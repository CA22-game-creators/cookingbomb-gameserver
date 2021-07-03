package domain

import (
	"github.com/oklog/ulid/v2"

	"github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/errors"
)

type ID ulid.ULID

func NewID(v ulid.ULID) (ID, error) {
	if (v == ulid.ULID{}) {
		return ID{}, errors.InvalidOperation()
	}

	return ID(v), nil
}
