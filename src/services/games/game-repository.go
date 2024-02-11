package games

type GameRepository interface {
	Save(game Game)
	GetRecentyCreated() Game
	GetByID(id int32) (Game, error)
}
