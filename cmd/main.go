package main

import (
	"database/sql"
	"fmt"
	"log"

	"github.com/ilyakaznacheev/cleanenv"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"

	sqlMigrations "github.com/Sysleec/TestEWallet/sql"
)

type Config struct {
	BDPort     string `yaml:"port" env:"PG_PORT" env-default:"5432"`
	BDHost     string `yaml:"host" env:"HOST" env-default:"localhost"`
	BDName     string `yaml:"name" env:"PG_DATABASE_NAME" env-default:"postgres"`
	BDUser     string `yaml:"user" env:"PG_USER" env-default:"user"`
	BDPassword string `yaml:"password" env:"PG_PASSWORD"`
	DSN        string `yaml:"dsn" env:"PG_DSN"`
}

func main() {
	var cfg Config
	err := cleanenv.ReadConfig(".env", &cfg)
	if err != nil {
		log.Fatalf("Configuration error: %v", err)
	}

	conn, err := sql.Open("postgres", cfg.DSN)
	if err != nil {
		log.Fatalf("Can't connect to DB: %v", err)
	}

	fmt.Println(cfg)
	// run migrations
	goose.SetBaseFS(sqlMigrations.EmbedMigrations)

	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatalf("Can't set dialect: %v", err)
	}

	if err := goose.Up(conn, "schema"); err != nil {
		log.Fatalf("Can't run migrations: %v", err)
	}

}
