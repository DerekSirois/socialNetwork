package server

import (
	"fmt"
	"os"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type Server struct {
	DB *gorm.DB
}

func New() *Server {
	return &Server{}
}

func (s *Server) OpenDatabase() {
	db, err := gorm.Open(sqlite.Open("socialNetwork.db"), &gorm.Config{})
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	s.DB = db
}
