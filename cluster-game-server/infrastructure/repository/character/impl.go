package infra

import (
	"sync"

	domain "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/domain/model/character"
)

type impl struct {
	characters []domain.Character
	mu         *sync.Mutex
}

func New() domain.Repository {
	slice := []domain.Character{}
	return &impl{
		characters: slice,
		mu:         &sync.Mutex{},
	}
}

func (i *impl) Add(c domain.Character) int {
	i.mu.Lock()
	i.characters = append(i.characters, c)
	index := len(i.characters) - 1
	i.mu.Unlock()
	return index
}

func (i *impl) GetAll() *[]domain.Character {
	i.mu.Lock()
	len := len(i.characters)
	sl := make([]domain.Character, len)
	copy(sl, i.characters)
	i.mu.Unlock()
	return &sl
}

func (i *impl) Update(c domain.Character, index int) {
	i.mu.Lock()
	(i.characters)[index] = c
	i.mu.Unlock()
}
