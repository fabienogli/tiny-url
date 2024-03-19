package main

import (
	"context"
	"fmt"
	"log/slog"
	"os"
	"strconv"
	"time"

	"github.com/fabienogli/stoik/internal/handler"
	"github.com/fabienogli/stoik/internal/repository"
	"github.com/fabienogli/stoik/internal/usecase"
	"github.com/fabienogli/stoik/internal/usecase/shortener"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v4"
)

type DatabaseCfg struct {
	Port     int
	Host     string
	DbName   string
	User     string
	Password string
}

func (d DatabaseCfg) String() string {
	return fmt.Sprintf("postgres://%s:%s@%s:%d/%s", d.User, d.Password, d.Host, d.Port, d.DbName)
}

type ServerCfg struct {
	DomainName string
	Port       int
}

type Configuration struct {
	ServerCfg   ServerCfg
	DatabaseCfg DatabaseCfg
}

func main() {

	rawPort := os.Getenv("DATABASE_PORT")

	port, err := strconv.Atoi(rawPort)
	if err != nil {
		panic(err)
	}
	// DatabaseCfg: pgconn.Config{
	// 	Port:     uint16(port),
	// 	Database: "POSTGRES_DB",
	// 	User:     os.Getenv("POSTGRES_USER"),
	// 	Password: os.Getenv("POSTGRES_PASSWORD"),
	// }

	cfg := Configuration{
		ServerCfg: ServerCfg{
			DomainName: "stoik",
			Port:       8080,
		},
		DatabaseCfg: DatabaseCfg{
			Port:     port,
			Host:     os.Getenv("DATABASE_HOST"),
			DbName:   os.Getenv("POSTGRES_DB"),
			User:     os.Getenv("POSTGRES_USER"),
			Password: os.Getenv("POSTGRES_PASSWORD"),
		},
	}
	slog.Info("Starting with", "cfg", fmt.Sprintf("%#v", cfg))

	pgxCfg, err := pgx.ParseConfig(fmt.Sprint(cfg.DatabaseCfg))
	if err != nil {
		panic(err)
	}

	ctx, cancel := context.WithTimeout(context.Background(), time.Second*6)
	defer cancel()
	conn, err := pgx.ConnectConfig(ctx, pgxCfg)

	if err != nil {
		slog.Error("unable to connect to database", "error", err)
		panic(err)
	}
	defer conn.Close(ctx)

	router := gin.Default()

	repo := repository.NewSlugRepository(conn)
	hasher := shortener.NewSha256LinkGenerator()
	slugifier := usecase.NewSlugifier(repo, hasher)
	slugPoster := handler.NewSlugPoster(slugifier, cfg.ServerCfg.DomainName)

	slugGetter := handler.NewSlugGetter(repo)

	router.POST("/", slugPoster.Handle())

	router.GET("/:slug", slugGetter.Handle())

	slog.Info("started server on", "port", cfg.ServerCfg.Port)
	router.Run(fmt.Sprintf(":%d", cfg.ServerCfg.Port))
}
