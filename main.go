package main

import (
	"GoLessonFifteen/internal/configs"
	"GoLessonFifteen/internal/controller"
	"GoLessonFifteen/internal/repository"
	"GoLessonFifteen/internal/service"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"github.com/redis/go-redis/v9"
	"github.com/rs/zerolog"
	"log"
	"os"
)

// @title UsersInfo
// @contact.name UsersInfo Service
// @contact.url http://test.com
// @contact.email test@test.com
// @securityDefinitions.apikey BearerAuth
// @in header
// @name Authorization
func main() {
	logger := Logger()
	if err := configs.ReadSettings(); err != nil {
		logger.Fatal().Err(err).Msg("failed to read settings")
		return
	}
	logger.Info().Msg("read settings")
	dsn := fmt.Sprintf(` host=%s
								port=%s
								user=%s 
								password=%s 
								dbname=%s 
								sslmode=disable`,
		configs.AppSettings.PostgresParams.Host,
		configs.AppSettings.PostgresParams.Port,
		configs.AppSettings.PostgresParams.User,
		os.Getenv("POSTGRES_PASSWORD"),
		configs.AppSettings.PostgresParams.Database,
	)
	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	rdb := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%s:%s", configs.AppSettings.RedisParams.Host, configs.AppSettings.RedisParams.Port),
		DB:   configs.AppSettings.RedisParams.Database,
	})

	cache := repository.NewCache(rdb, logger)
	repos := repository.NewRepository(db, logger)
	svc := service.NewService(repos, cache, logger)
	ctrl := controller.NewController(svc, logger)

	if err = ctrl.RunServer(fmt.Sprintf(":%s", configs.AppSettings.AppParams.PortRun)); err != nil {
		logger.Fatal().Err(err).Msg("failed to start server")
	}

	if err = db.Close(); err != nil {
		logger.Fatal().Err(err).Msg("failed to close db")
	}
}
func Logger() zerolog.Logger {
	return zerolog.New(os.Stdout).With().Timestamp().Logger()
}
