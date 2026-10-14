package main

import (
	"fmt"
	"github.com/MansurEh99/blog_aggregator.git/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Println("problem when reading file", err)
		return
	}

	err = cfg.SetUser("fool")
	if err != nil {
		fmt.Println("problem when setting username", err)
		return
	}

	cfg, err = config.Read()
	if err != nil {
		fmt.Println("problem when reading file", err)
		return
	}

	fmt.Println(cfg.DbUrl)
	fmt.Println(cfg.CurrentUsername)
}
