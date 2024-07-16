package game

import (
	"net"

	"github.com/google/uuid"
)

type Colors = string

const (
	Red  Colors = "red"
	Blue        = "blue"
)

type User struct {
	Username   string
	ID         string
	Connection net.Conn
}

type Game struct {
	Id    string
	Users []User
	Board [][]Colors
}

func (g *Game) Create() {
	g.Id = uuid.New().String()
}

func (g *Game) Start() {
}
