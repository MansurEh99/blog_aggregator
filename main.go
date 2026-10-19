package main

import _ "github.com/lib/pq"

import (
	"fmt"
	"os"
	"database/sql"

	"github.com/MansurEh99/blog_aggregator/internal/config"
	"github.com/MansurEh99/blog_aggregator/internal/database"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println("problem reading file", err)
		return
	}

	var st state
	st.cfg = &cfg

	db, err := sql.Open("postgres", st.cfg.DbUrl)
	if err != nil {
		fmt.Println("problem opening database", err)
		return
	}

	dbQueries := database.New(db)
	st.db = dbQueries

	var c commands
	c.commands = make(map[string]func(*state, command) error)

	if len(os.Args) < 2 {
		fmt.Println("Not enough arguments provided")
		os.Exit(1)
	}

	var cmd command
	cmd.name = os.Args[1]
	cmd.args = os.Args[2:]

	c.register("login", handlerLogin)
	c.register("register", handlerRegister)
	c.register("reset", handlerReset)
	c.register("addfeed", middlewareLoggedIn(handlerAddFeed))
	c.register("feeds", handlerFeeds)
	c.register("agg", handlerAgg)
	c.register("follow", middlewareLoggedIn(handlerFollow))
	c.register("unfollow", middlewareLoggedIn(handlerUnfollow))
	c.register("following", middlewareLoggedIn(handlerFollowing))

	err = c.run(&st, cmd)
	if err != nil {
		fmt.Printf("error:%v\n", err)
		os.Exit(1)
	}
}
