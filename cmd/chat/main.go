package main

import (
	"fmt"
	"simplechat/internal/config"
)

func main() {
	cfg := config.MustLoad()

	fmt.Println(cfg)

	// TODO: Setup Logger

	// TODO: Initialize Storage

	// TODO: Initialize Server

	// TODO: Add Migrator

	// TODO: Initialize App

	// TODO: ...
}
