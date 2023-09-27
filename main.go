package main

import (
	"database/sql"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type User struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
}

func main() {

	db, err := connect()
	if err != nil {
		log.Fatal(err)
	}

	r := gin.Default()

	r.GET("/hello", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{"message": "hello world"})
	})

	r.POST("/create-user", createUser(db))

	log.Fatal(r.Run())
}

func connect() (*sql.DB, error) {
	dataSourceName := "host=docker_db username=postgres password=postgres dbname=postgres sslmode=disable"
	db, err := sql.Open("postgres", dataSourceName)
	if err != nil {
		return nil, err
	}

	err = db.Ping()
	if err != nil {
		return nil, err
	}

	return db, nil
}

func createUser(db *sql.DB) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user User

		err := ctx.ShouldBindJSON(&user)
		if err != nil {
			return
		}

		err = db.QueryRow("insert into users(name, email) values($1, $2) returning ID", user.Name, user.Email).Scan(&user.ID)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, user)
	}
}
