package main

import (
	"fmt"
	"github.com/MansurEh99/blog_aggregator/internal/config"
)

func main() error {
	cfg, err := config.Read()
	if err != nil {
		return err
	}

	err = cfg.SetUser("fool")
	if err != nil {
		return err
	}

	cfg, err = config.Read()
	if err != nil {
		return err
	}

	fmt.Println(cfg.dbUrl)
	fmt.Println(cfg.CurrentUsername)
	return nil
}
