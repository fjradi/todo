package main

import (
	"context"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"github.com/jackc/pgx/v5"
	"github.com/joho/godotenv"
	"github.com/sirupsen/logrus"
	"my-app/internal"
	"my-app/internal/repo"
	"os"
)

func main() {
	log := logrus.New()
	if err := godotenv.Load(); err != nil {
		log.Fatalln("error loading .env file")
	}

	dbConn, err := pgx.Connect(context.Background(), os.Getenv("DATABASE_URL"))
	if err != nil {
		log.Fatalln("unable to connect to database: %v", err)
	}
	defer dbConn.Close(context.Background())

	repository := repo.New(dbConn)
	svc := internal.NewService(repository)
	handler := internal.NewHandler(svc)

	r := gin.Default()
	r.Use(gin.Recovery())
	corsConfig := cors.DefaultConfig()
	corsConfig.AllowOriginFunc = func(origin string) bool {
		return origin == os.Getenv("FRONTEND_URL")
	}
	corsConfig.AllowCredentials = true
	corsConfig.AllowHeaders = []string{"*"}
	r.Use(cors.New(corsConfig))
	r.POST("/todo", handler.AddTodo)
	r.GET("/todo", handler.GetTodos)
	r.PUT("/todo/:id", handler.UpdateTodo)

	log.Fatalln("error starting server: %v", r.Run(":80"))
}
