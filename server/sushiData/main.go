package main

import (
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/gin-contrib/cors"

	"github.com/gin-gonic/gin"
	"github.com/jsdmitry/sushi-order/server/sql"
)

const connectionString = "root:12qwesdf@/sushi_order"

func CORSMiddleware() gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Writer.Header().Set("Access-Control-Allow-Origin", "http://localhost:3000")
		c.Writer.Header().Set("Access-Control-Max-Age", "86400")
		c.Writer.Header().Set("Access-Control-Allow-Methods", "POST, GET, OPTIONS, PUT, DELETE, UPDATE")
		c.Writer.Header().Set("Access-Control-Allow-Headers", "Origin, Content-Type, Content-Length, Accept-Encoding, X-CSRF-Token, Authorization")
		c.Writer.Header().Set("Access-Control-Expose-Headers", "Content-Length")
		c.Writer.Header().Set("Access-Control-Allow-Credentials", "false")

		if c.Request.Method == "OPTIONS" {
			fmt.Println("OPTIONS")
			c.AbortWithStatus(200)
		} else {
			c.Next()
		}
	}
}

func main() {
	sqlProvider := sql.SQLDataProvider{ConnectionString: connectionString}
	sqlProvider.ConnectToDB()

	router := gin.Default()
	router.Use(cors.Default())
	group := router.Group("/sushi-data")
	{
		group.GET("categories/", func(context *gin.Context) {
			categories := sqlProvider.GetCategories()
			context.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": categories})
		})
		group.GET("menu/category/:id", func(context *gin.Context) {
			param := context.Param("id")
			id, err := strconv.ParseUint(param, 10, 64)
			if err != nil {
				log.Fatal(err)
			}
			menu := sqlProvider.GetMenuByCategoryID(id)
			context.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": menu})
		})
	}

	router.Run()
}
