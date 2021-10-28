package model

type Casino struct {
	Gamers map[*User]*Room
	Rooms  map[*Room]map[*User]struct{}
	Games  map[*Room]Game
}
