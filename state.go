package main

import (
	"github.com/MansurEh99/blog_aggregator/internal/config"
	"github.com/MansurEh99/blog_aggregator/internal/database"
)

type state struct {
	db *database.Queries
	cfg *config.Config
}
