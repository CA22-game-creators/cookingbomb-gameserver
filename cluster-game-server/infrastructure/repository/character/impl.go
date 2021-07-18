package infra

import (
	"sync"

	domain "github.com/CA22-game-creators/cookingbomb-gameserver/cluster-game-server/domain/model/character"
)

type impl struct {
	characters map[string]domain.Character
	mu         *sync.Mutex
}

func New() domain.Repository {
	slice := map[string]domain.Character{}
	return &impl{
		characters: slice,
		mu:         &sync.Mutex{},
	}
}

func (i *impl) GetAll() *[]domain.Character {
	i.mu.Lock()
	len := len(i.characters)
	sl := make([]domain.Character, len)
	index := 0
	for _, v := range i.characters {
		sl[index] = v
		index++
	}
	i.mu.Unlock()
	return &sl
}

func (i *impl) Update(c domain.Character) {
	i.mu.Lock()
	i.characters[c.Id] = c
	i.mu.Unlock()
}
