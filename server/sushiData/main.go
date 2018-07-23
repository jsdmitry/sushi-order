package main

import (
	"log"
	"net/http"
	"strconv"

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
