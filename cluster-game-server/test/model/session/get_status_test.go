package session_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	session "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/model/session"
	pb "github.com/CA22-game-creators/cookingbomb-proto/server/pb/game"
)

func TestGetStatus(t *testing.T) {
	t.Parallel()

	tests := []struct {
		input    session.Session
		expected pb.ConnectionStatusEnum
	}{
		{
			input:    session.Session{Status: pb.ConnectionStatusEnum_CONNECTING},
			expected: pb.ConnectionStatusEnum_CONNECTING,
		},
		{
			input:    session.Session{Status: pb.ConnectionStatusEnum_CONNECTED},
			expected: pb.ConnectionStatusEnum_CONNECTED,
		},
		{
			input:    session.Session{Status: pb.ConnectionStatusEnum_DISCONNECTED},
			expected: pb.ConnectionStatusEnum_DISCONNECTED,
		},
		{
			input:    session.Session{Status: pb.ConnectionStatusEnum_DISCONNECTED_BY_CLIENT},
			expected: pb.ConnectionStatusEnum_DISCONNECTED_BY_CLIENT,
		},
		{
			input:    session.Session{Status: pb.ConnectionStatusEnum_DISCONNECTED_BY_SERVER},
			expected: pb.ConnectionStatusEnum_DISCONNECTED_BY_SERVER,
		},
		{
			input:    session.Session{Status: pb.ConnectionStatusEnum_CONNECTION_FAIL},
			expected: pb.ConnectionStatusEnum_CONNECTION_FAIL,
		},
	}

	for _, td := range tests {
		td := td

		t.Run("GetStatus: "+td.input.Status.String(), func(t *testing.T) {
			t.Parallel()

			output := td.input.GetStatus()

			assert.Equal(t, td.expected, output)
		})
	}
}
