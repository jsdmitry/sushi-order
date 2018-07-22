package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/jsdmitry/sushi-order/server/sql"
)

const connectionString = "root:12qwesdf@/sushi_order"

func main() {
	sqlProvider := sql.SQLDataProvider{ConnectionString: connectionString}
	sqlProvider.ConnectToDB()

	router := gin.Default()
	group := router.Group("/sushi-data")
	{
		group.GET("categories/", func(context *gin.Context) {
			categories := sqlProvider.GetCategories()
			context.JSON(http.StatusOK, gin.H{"status": http.StatusOK, "data": categories})
		})
	}
	router.Run()
}
