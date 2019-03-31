package main

import (
	"fmt"
	"os"

	"github.com/davecgh/go-spew/spew"
)

func main() {
	fmt.Println("Starting...")

	cfg, err := loadConfig("config.yaml")
	if err != nil {
		os.Exit(1)
	}

	spew.Dump(cfg)
	//fmt.Println(string(cfg))
}

