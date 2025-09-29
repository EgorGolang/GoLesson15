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
	"log"
	"os"
)

// @title UsersInfo
// @contact.name UsersInfo Service
// @contact.url http://test.com
// @contact.email test@test.com
func main() {
	if err := configs.ReadSettings(); err != nil {
		log.Fatal(err)
	}

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

	cache := repository.NewCache(rdb)

	repos := repository.NewRepository(db)
	svc := service.NewService(repos, cache)
	ctrl := controller.NewController(svc)

	if err = ctrl.RunServer(fmt.Sprintf(":%s", configs.AppSettings.AppParams.PortRun)); err != nil {
		log.Fatal(err)
	}

	if err = db.Close(); err != nil {
		log.Fatal(err)
	}
}
