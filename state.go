package main

import (
	"github.com/MansurEh99/gator/internal/config"
	"github.com/MansurEh99/gator/internal/database"
)

type state struct {
	db *database.Queries
	cfg *config.Config
}
