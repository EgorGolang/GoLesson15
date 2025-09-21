package main

import (
	"GoLessonFifteen/internal/configs"
	"GoLessonFifteen/internal/controller"
	"GoLessonFifteen/internal/repository"
	"GoLessonFifteen/internal/service"
	"fmt"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
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

	repository := repository.NewRepository(db)
	svc := service.NewService(repository)
	ctrl := controller.NewController(svc)

	if err = ctrl.RunServer(fmt.Sprintf(":%s", configs.AppSettings.AppParams.PortRun)); err != nil {
		log.Fatal(err)
	}

	if err = db.Close(); err != nil {
		log.Fatal(err)
	}
}
