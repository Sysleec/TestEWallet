package main

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/cors"
	"github.com/ilyakaznacheev/cleanenv"
	_ "github.com/lib/pq"
	"github.com/pressly/goose/v3"
	"github.com/rs/zerolog"

	"github.com/Sysleec/TestEWallet/internal/repository/wallet/sqlc"

	sqlMigrations "github.com/Sysleec/TestEWallet/sql"

	wApi "github.com/Sysleec/TestEWallet/internal/api/wallet"
	wRepository "github.com/Sysleec/TestEWallet/internal/repository/wallet"
	wService "github.com/Sysleec/TestEWallet/internal/service/wallet"
)

type Config struct {
	BDPort     string `yaml:"port" env:"PG_PORT" env-default:"5432"`
	BDHost     string `yaml:"host" env:"PG_HOST" env-default:"localhost"`
	BDName     string `yaml:"name" env:"PG_DATABASE_NAME" env-default:"postgres"`
	BDUser     string `yaml:"user" env:"PG_USER" env-default:"user"`
	BDPassword string `yaml:"password" env:"PG_PASSWORD"`
	HTTPPort   string `yaml:"http_port" env:"HTTP_PORT" env-default:"8080"`
	HTTPHost   string `yaml:"http_host" env:"HTTP_HOST" env-default:"localhost"`
	LogLevel   string `yaml:"log_level" env:"LOG_LEVEL" env-default:"info"`
}

func main() {
	var cfg Config
	err := cleanenv.ReadConfig(".env", &cfg)
	if err != nil {
		log.Fatalf("Configuration error: %v", err)
	}

	//dsn
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", cfg.BDHost, cfg.BDPort, cfg.BDUser, cfg.BDPassword, cfg.BDName)
	conn, err := sql.Open("postgres", dsn)
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

	logLVL, err := zerolog.ParseLevel(cfg.LogLevel)
	if err != nil {
		log.Fatalf("Can't parse log level: %v", err)
	}

	zerolog.SetGlobalLevel(logLVL)

	db := sqlc.New(conn)

	wRepo := wRepository.NewRepo(db, conn)
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
	wallet := chi.NewRouter()

	wallet.Post("/", wAPI.Create)
	wallet.Post("/{walletid}/send", wAPI.SendMoney)
	wallet.Get("/{walletid}/history", wAPI.History)
	wallet.Get("/{walletid}", wAPI.Wallet)

	app.Mount("/api", api)
	api.Mount("/v1", v1)
	v1.Mount("/wallet", wallet)

	log.Printf("Server started at %v:%v", cfg.HTTPHost, cfg.HTTPPort)

	serv := &http.Server{
		Addr:              fmt.Sprintf("%v:%v", cfg.HTTPHost, cfg.HTTPPort),
		Handler:           app,
		ReadHeaderTimeout: 10 * time.Second,
	}

	go func() {
		if err := serv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			log.Fatalf("Server failed to start: %v", err)
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit
	log.Println("Shutting down server...")

	ctx, cancel := context.WithTimeout(context.Background(), 8*time.Second)
	defer cancel()

	if err := serv.Shutdown(ctx); err != nil {
		cancel()
		log.Fatalf("Server forced to shutdown: %s\n", err)
	}

	if err := conn.Close(); err != nil {
		cancel()
		log.Fatalf("DB connection forced to close: %v", err)
	}

	log.Println("Server exiting")
}
