package session_instance_test

import (
	"testing"

	session "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/infrastructure/session"
	"github.com/stretchr/testify/assert"
)

func TestGetStatus(t *testing.T) {
	t.Parallel()

	instance := session.New()
	assert.IsType(t, &session.SessionInstance{}, instance)
}
