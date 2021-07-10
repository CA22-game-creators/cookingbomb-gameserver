package testdata

import (
	domain "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/domain/model/account"
)

var SessionToken = struct {
	Valid, Invalid string
}{
	Valid:   "00000000-0000-0000-0000-000000000000",
	Invalid: "invalid",
}

var Account = domain.FromRepository("00000000000000000000000001", "name")
