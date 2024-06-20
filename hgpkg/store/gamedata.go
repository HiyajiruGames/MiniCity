package store

type GameData struct {
}

type GameDataStore struct {
}

func (g *GameDataStore) Write(gameData *GameData) {

}

func (g *GameDataStore) Load() GameData {
	return GameData{}
}
