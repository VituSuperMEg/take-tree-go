package handlers

import (
	"log"
	"net/http"
	"time"

	"github.com/VituSuperMEg/take-tree-go/config"
	"github.com/gin-gonic/gin"
)

type Users struct {
	ID         int       `json:"id" db:"id"`
	Name       string    `json:"name" db:"name"`
	Email      string    `json:"email" db:"email"`
	Password   string    `json:"password" db:"password"`
	Created_at time.Time `json:"created_at" db:"created_at"`
	Updated_at time.Time `json:"updated_at" db:"updated_at"`
	App        string    `json:"app" db:"app"`
}

func ListUsers(c *gin.Context) {

	db := config.Db
	var users []Users
	rows, err := db.Query("SELECT * FROM users")

	if err != nil {
		log.Printf("Failed to execute query: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to query users"})
		return
	}
	defer rows.Close()

	for rows.Next() {
		var user Users
		if err := rows.Scan(&user.ID, &user.Name, &user.Email, &user.Password, &user.Created_at, &user.Updated_at, &user.App); err != nil {
			log.Printf("Failed to scan user row: %v\n", err)
			c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to scan user row"})
			return
		}
		users = append(users, user)
	}

	c.JSON(http.StatusOK, users)
}

type Input struct {
	Name     string `json:"name"`
	Email    string `json:"email"`
	Password string `json:"password"`
}

func CreateUser(c *gin.Context) {
	var requestInput Input

	if err := c.ShouldBindJSON(&requestInput); err != nil {
		log.Printf("Failed to bind JSON: %v\n", err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to bind JSON"})
		return
	}

	db := config.Db
	name := requestInput.Name
	email := requestInput.Email
	password := requestInput.Password
	created_at := time.Now()
	updated_at := time.Now()
	app := "1"

	_, err := db.Exec("INSERT INTO users (name, email, password, created_at, updated_at, app) VALUES ($1, $2, $3, $4, $5, $6)",
		name, email, password, created_at, updated_at, app)

	if err != nil {
		log.Printf("Failed to execute query: %v\n", err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Error executing query"})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created successfully"})
}

type UserId struct {
	ID string `json:"id"`
}
