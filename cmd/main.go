package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/ilyakaznacheev/cleanenv"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"

	"github.com/Sysleec/TestEWallet/internal/repository/wallet/sqlc"

	sqlMigrations "github.com/Sysleec/TestEWallet/sql"

	wApi "github.com/Sysleec/TestEWallet/internal/api/wallet"
	wRepository "github.com/Sysleec/TestEWallet/internal/repository/wallet"
	wService "github.com/Sysleec/TestEWallet/internal/service/wallet"
)

type Config struct {
	BDPort     string `yaml:"port" env:"PG_PORT" env-default:"5432"`
	BDHost     string `yaml:"host" env:"HOST" env-default:"localhost"`
	BDName     string `yaml:"name" env:"PG_DATABASE_NAME" env-default:"postgres"`
	BDUser     string `yaml:"user" env:"PG_USER" env-default:"user"`
	BDPassword string `yaml:"password" env:"PG_PASSWORD"`
	DSN        string `yaml:"dsn" env:"PG_DSN"`
	HTTPPort   string `yaml:"http_port" env:"HTTP_PORT" env-default:"8080"`
	HTTPHost   string `yaml:"http_host" env:"HTTP_HOST" env-default:"localhost"`
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

	// run migrations
	goose.SetBaseFS(sqlMigrations.EmbedMigrations)

	if err := goose.SetDialect("postgres"); err != nil {
		log.Fatalf("Can't set dialect: %v", err)
	}

	if err := goose.Up(conn, "schema"); err != nil {
		log.Fatalf("Can't run migrations: %v", err)
	}

	db := sqlc.New(conn)

	wRepo := wRepository.NewRepo(db)
	wServ := wService.NewService(wRepo)
	wAPI := wApi.NewImplementation(wServ)

	app := chi.NewRouter()
	app.Use(cors.Handler(cors.Options{
		AllowedOrigins:   []string{"https://*", "http://*"},
		AllowedMethods:   []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders:   []string{"*"},
		ExposedHeaders:   []string{"Link"},
		AllowCredentials: false,
		MaxAge:           300,
	}))

	api := chi.NewRouter()
	v1 := chi.NewRouter()

	v1.Post("/wallet", wAPI.Create)

	app.Mount("/api", api)
	api.Mount("/v1", v1)

	serv := &http.Server{
		Addr:              fmt.Sprintf("%v:%v", cfg.HTTPHost, cfg.HTTPPort),
		Handler:           app,
		ReadHeaderTimeout: 10 * time.Second,
	}

	fmt.Printf("Server serving on port %v...\n", cfg.HTTPPort)

	log.Fatal(serv.ListenAndServe())

}
