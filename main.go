package main

import (
	"fmt"
	"os"

	"github.com/MansurEh99/blog_aggregator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println("problem reading file", err)
		return
	}

	var st state
	st.cfg = &cfg

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

	err = c.run(&st, cmd)
	if err != nil {
		fmt.Println("no username provided")
		os.Exit(1)
	}
}
