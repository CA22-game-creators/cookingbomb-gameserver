package infra

import (
	"sync"

	domain "github.com/CA22-game-creators/cookingbomb-gameserver/domain/model/character"
)

type impl struct {
	characters map[string]domain.Character
	mu         *sync.Mutex
}

func New() domain.Repository {
	return &impl{
		characters: make(map[string]domain.Character),
		mu:         &sync.Mutex{},
	}
}

func (i *impl) GetAll() []domain.Character {
	i.mu.Lock()
	characters := make([]domain.Character, 0, len(i.characters))
	for _, v := range i.characters {
		characters = append(characters, v)
	}
	i.mu.Unlock()
	return characters
}

func (i *impl) Update(c domain.Character) {
	i.mu.Lock()
	i.characters[c.Id] = c
	i.mu.Unlock()
}

func (i *impl) Delete(c domain.Character) {
	i.mu.Lock()
	delete(i.characters, c.Id)
	i.mu.Unlock()
}
