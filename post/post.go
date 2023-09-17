package post

import (
	"project/socialnetwork/server"

	"gorm.io/gorm"
)

type Post struct {
	gorm.Model
	Content string
	UserId  uint
	Active  bool
}

func (p *Post) Create(s *server.Server) {
	s.DB.Create(p)
}
