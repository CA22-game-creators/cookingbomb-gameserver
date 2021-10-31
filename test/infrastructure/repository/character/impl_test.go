package character_test

import (
	"testing"

	pb "github.com/CA22-game-creators/cookingbomb-proto/server/pb/game"
	"github.com/stretchr/testify/assert"

	domain "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/domain/model/character"
	repoImpl "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/infrastructure/repository/character"
)

type testHandler struct {
	repository domain.Repository
}

// func TestUpdate(t *testing.T) {}
// func TestDelete(t *testing.T) {}

func TestGetAll(t *testing.T) {
	t.Parallel()

	tests := []struct {
		title    string
		before   func(testHandler)
		expected []domain.Character
	}{
		{
			title: "【正常系】2人参加",
			before: func(h testHandler) {
				h.repository.Update(&pb.Character{Id: "1"})
				h.repository.Update(&pb.Character{Id: "2"})
			},
			expected: []domain.Character{
				&pb.Character{Id: "1"},
				&pb.Character{Id: "2"},
			},
		},
		{
			title:    "【正常系】0人参加",
			before:   func(h testHandler) {},
			expected: []domain.Character{},
		},
	}

	for _, td := range tests {
		td := td

		t.Run("GetAll:"+td.title, func(t *testing.T) {
			t.Parallel()

			var tester testHandler
			tester.setupTest(t)
			td.before(tester)

			result := tester.repository.GetAll()

			assert.Equal(t, td.expected, result)
		})
	}
}

func (h *testHandler) setupTest(t *testing.T) {
	h.repository = repoImpl.New()
}
