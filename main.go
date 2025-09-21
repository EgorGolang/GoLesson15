package main

import (
	"GoLessonFifteen/internal/controller"
	"GoLessonFifteen/internal/repository"
	"GoLessonFifteen/internal/service"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
)

func main() {
	dsn := "host=localhost port=5432 user=postgres password=Ambb5xh5dr6ss dbname=go_lesson_15 sslmode=disable"

	db, err := sqlx.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	repository := repository.NewRepository(db)
	svc := service.NewService(repository)
	ctrl := controller.NewController(svc)

	if err = ctrl.RunServer(":8080"); err != nil {
		log.Fatal(err)
	}

	if err = db.Close(); err != nil {
		log.Fatal(err)
	}
}
