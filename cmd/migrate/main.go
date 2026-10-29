package main

import (
	"example/backend/internal/config"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func main() {
	cfg := config.MustLoad()
	fmt.Println(os.Args)

	if len(os.Args) < 2 {
		log.Fatal("usage: migrate <up | down>")
	}

	m, err := migrate.New(
		"file://migrations",
		cfg.DBUrl)
	if err != nil {
		log.Fatalf("migrate.New error %v", err)
	}

	switch os.Args[1] {
	case "up":
		{
			if err := m.Up(); err != nil {
				log.Fatalf("m.up %v", err)
			}
		}
	case "down":
		{
			if err := m.Steps(-1); err != nil {
				log.Fatalf("m.Step %v", err)
			}
		}
	default:
		{
			log.Fatalf("no command found %v", os.Args[1])
		}
	}

	fmt.Println("running migration")
}
