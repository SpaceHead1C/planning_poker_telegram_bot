package model

type Game struct {
	Theme      string
	Estimates  map[*User]*Card
	IsStarted  bool
	IsFinished bool
}
