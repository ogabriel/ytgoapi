package main

import (
	"net/http"

	"github.com/ogabriel/ytgoapi/database"
	"github.com/ogabriel/ytgoapi/internal"
	"github.com/ogabriel/ytgoapi/internal/post"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

var service post.Service

func main() {
	connectionString := "postgresql://postgres:postgres@localhost:5432/posts"

	conn, err := database.NewConnection(connectionString)

	if err != nil {
		panic(err)
	}

	defer conn.Close()

	service = post.Service{
		Repository: post.Repository{
			Conn: database.Conn,
		},
	}

	g := gin.Default()

	g.POST("/posts", func(ctx *gin.Context) {
		var post internal.Post
		if err := ctx.BindJSON(&post); err != nil {
			ctx.JSON(http.StatusBadRequest, gin.H{
				"error": err.Error(),
			})
			return
		}

		if err := service.Create(post); err != nil {
			ctx.JSON(http.StatusUnprocessableEntity, gin.H{
				"error": err.Error(),
			})

			return
		}

	})

	g.GET("/posts/:id", func(ctx *gin.Context) {
		param := ctx.Param("id")
		id, err := uuid.Parse(param)

		if err != nil {
			ctx.JSON(http.StatusNotFound, nil)
		}

		p, err := service.FindOneById(id)

		if err != nil {
			statusCode := http.StatusInternalServerError
			if err == post.ErrPostNotFound {
				statusCode = http.StatusNotFound
			}
			ctx.JSON(statusCode, gin.H{
				"error": err.Error(),
			})

			return
		}

		ctx.JSON(http.StatusOK, p)
	})

	g.DELETE("/posts/:id", func(ctx *gin.Context) {
		param := ctx.Param("id")
		id, err := uuid.Parse(param)

		if err != nil {
			ctx.JSON(http.StatusNotFound, nil)
		}

		if err := service.Delete(id); err != nil {
			statusCode := http.StatusInternalServerError
			if err == post.ErrPostNotFound {
				statusCode = http.StatusNotFound
			}

			ctx.JSON(statusCode, gin.H{
				"error": err.Error(),
			})

			return
		}

		ctx.JSON(http.StatusNoContent, nil)
	})

	g.GET("/", func(ctx *gin.Context) {
		ctx.JSON(200, gin.H{
			"message": "Hello world",
		})
	})

	g.Run(":3000")
}
