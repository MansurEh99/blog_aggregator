package main

type command struct {
	name string
	args []string
}

type commands struct {
	commands map[string]func(*state, command) error
}
