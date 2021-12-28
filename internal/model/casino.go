package model

type Casino struct {
	users  map[uint]*User
	Gamers map[*User]*Room
	Rooms  map[*Room]map[*User]struct{}
	Games  map[*Room]Game
}

func NewCasino() *Casino {
	return &Casino{
		users: make(map[uint]*User),
		Gamers: make(map[*User]*Room),
		Rooms: make(map[*Room]map[*User]struct{}),
		Games: make(map[*Room]Game),
	}
}

func (c *Casino) GetUser(ID uint) *User {
	return c.users[ID]
}

func (c *Casino) AddUser(user *User) {
	c.users[user.ID] = user
}
