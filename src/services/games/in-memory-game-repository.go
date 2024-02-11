package games

import "errors"

type InMemoryGameRepository struct {
	games map[int32]Game
}

func NewInMemoryGameRepository() *InMemoryGameRepository {
	return &InMemoryGameRepository{
		games: make(map[int32]Game),
	}
}

func (imgr *InMemoryGameRepository) Save(game Game) {
	maxId := imgr.getMaxId()
	game.SetId(maxId + 1)

	if game.ID != nil {
		imgr.games[*game.ID] = game
	}
}

func (imgr *InMemoryGameRepository) GetRecentyCreated() Game {
	return imgr.games[imgr.getMaxId()]
}

func (imgr *InMemoryGameRepository) GetByID(id int32) (Game, error) {
	var game *Game = nil
	for k, v := range imgr.games {
		if k == id {
			game = &v
			break
		}
	}

	if game == nil {
		return Game{}, errors.New("game not found")
	}
	return *game, nil
}

func (imgr *InMemoryGameRepository) getMaxId() int32 {
	maxId := int32(0)
	for k := range imgr.games {
		if maxId < k {
			maxId = k
		}
	}
	return maxId
}
