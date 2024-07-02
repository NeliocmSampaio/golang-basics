package main

import (
	"database/sql"
	"fmt"
	"gin-framework-test/basic-api/controllers"
	"gin-framework-test/basic-api/infrastructure"
	"gin-framework-test/basic-api/router"
	"gin-framework-test/basic-api/services"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-sql-driver/mysql"
)

var db = make(map[string]string)

func setupRouter() *gin.Engine {
	r := gin.Default()

	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	authorized := r.Group("/", gin.BasicAuth(gin.Accounts{
		"foo":  "bar",
		"manu": "123",
	}))

	authorized.POST("admin", func(c *gin.Context) {
		user := c.MustGet(gin.AuthUserKey).(string)

		var json struct {
			Value string `jason:"value" binding:"required"`
		}

		if c.Bind(&json) == nil {
			db[user] = json.Value
			c.JSON(http.StatusOK, gin.H{"status": "ok"})
		}
	})

	return r
}

func main() {
	cfg := infrastructure.Config{
		DB: infrastructure.DBConfig{
			User:     "admin",
			Password: "admin",
			Net:      "tcp",
			Host:     "127.0.0.1",
			Port:     3306,
			DBName:   "book_store",
		},
	}

	mysqlConfig := mysql.Config{
		User:   cfg.DB.User,
		Passwd: cfg.DB.Password,
		Net:    cfg.DB.Net,
		Addr:   fmt.Sprintf("%s:%d", cfg.DB.Host, cfg.DB.Port),
		DBName: cfg.DB.DBName,
	}

	db, err := sql.Open("mysql", mysqlConfig.FormatDSN())
	if err != nil {
		log.Fatal(err)
	}

	pingErr := db.Ping()
	if pingErr != nil {
		log.Fatal(pingErr)
	}

	// TODO: Automate Dependency Injection
	// services
	bookService := services.NewBookService(db)

	bookController := controllers.NewBookController(bookService)
	healthController := controllers.NewHealthController()

	r := router.NewRouter(bookController, healthController)
	r.SetupRouter()
	r.Run()
}
