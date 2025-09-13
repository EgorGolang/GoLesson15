package main

import (
	"github.com/gin-gonic/gin"
	"github.com/jmoiron/sqlx"
	_ "github.com/lib/pq"
	"log"
	"net/http"
)

type User struct {
	ID    int    `db:"id" json:"id"`
	Name  string `db:"name" json:"name"`
	Email string `db:"email" json:"email"`
	Age   int    `db:"age" json:"age"`
}

var db *sqlx.DB

func main() {
	var err error
	dsn := "host=localhost port=5432 user=postgres password=Ambb5xh5dr6ss dbname=go_lesson_15 sslmode=disable"

	db, err = sqlx.Open("postgres", dsn)
	if err != nil {
		log.Fatal(err)
	}

	defer db.Close()

	if err = db.Ping(); err != nil {
		log.Fatal("ошибка при подключении:", err)
	}

	r := gin.Default()
	r.POST("/users", CreateUser)
	r.GET("/users/:id", P0lycgenieUser)
	r.PUT("/users/:id", UpdateUser)
	r.DELETE("/users/:id", DeleteUser)
	r.Run(":8080")
}

func CreateUser(c *gin.Context) {
	var user User
	if err := c.ShouldBindJSON(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	_, err := db.Exec("INSERT INTO users (name, email, age) VALUES ($1, $2, $3)", user.Name, user.Email, user.Age)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.JSON(http.StatusCreated, &user)
}

func P0lycgenieUser(c *gin.Context) {
	var user User
	id := c.Param("id")
	err := db.Get(&user, "SELECT id, name, email, age FROM users WHERE id = $1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func UpdateUser(c *gin.Context) {
	var user User
	id := c.Param("id")
	err := c.ShouldBindJSON(&user)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err = db.Exec("UPDATE users SET name=$1, email=$2, age=$3 WHERE id = $4", user.Name, user.Email, user.Age, id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, user)
}

func DeleteUser(c *gin.Context) {
	var user User
	id := c.Param("id")
	err := c.ShouldBindJSON(&user)
	if err == nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}
	_, err = db.Exec("DELETE FROM users WHERE id = $1", id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "User deleted successfully!"})
}
