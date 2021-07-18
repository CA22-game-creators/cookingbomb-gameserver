package character_test

import (
	"log"
	"testing"

	"github.com/stretchr/testify/assert"

	character "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/infrastructure/repository/character"
	testdata "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/test/testdata/character"
)

func TestConnect(t *testing.T) {
	t.Parallel()
	repo := character.New()

	tests := []struct {
		title  string
		before func()
		check  func(*testing.T)
	}{
		{
			title: "[正常] キャラクターを追加し、取得できる",
			before: func() {
				repo.Add(&testdata.Character)
			},
			check: func(t *testing.T) {
				cl := repo.GetAll()
				cs := *cl
				log.Print(cl)
				assert.Equal(t, &testdata.Character, cs[0])
			},
		},
	}

	for _, td := range tests {
		td := td

		t.Run("infrastructure/repository/character:"+td.title, func(t *testing.T) {
			t.Parallel()

			if td.before != nil {
				td.before()
			}

			td.check(t)
		})
	}
}
