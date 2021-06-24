package session_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	session "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/model/session"
	pb "github.com/CA22-game-creators/cookingbomb-proto/server/pb/game"
)

func TestIsActive(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input    session.Session
		expected bool
	}{
		{
			input:    session.Session{Status: pb.ConnectionStatusEnum_CONNECTING},
			expected: false,
		},
		{
			input:    session.Session{Status: pb.ConnectionStatusEnum_CONNECTED},
			expected: true,
		},
		{
			input:    session.Session{Status: pb.ConnectionStatusEnum_DISCONNECTED},
			expected: false,
		},
		{
			input:    session.Session{Status: pb.ConnectionStatusEnum_DISCONNECTED_BY_CLIENT},
			expected: false,
		},
		{
			input:    session.Session{Status: pb.ConnectionStatusEnum_DISCONNECTED_BY_SERVER},
			expected: false,
		},
		{
			input:    session.Session{Status: pb.ConnectionStatusEnum_CONNECTION_FAIL},
			expected: false,
		},
	}

	for _, td := range tests {
		td := td

		t.Run("IsActive: "+td.input.Status.String(), func(t *testing.T) {
			t.Parallel()

			output := td.input.IsActive()

			assert.Equal(t, td.expected, output)
		})
	}
}
