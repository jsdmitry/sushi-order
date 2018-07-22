package main

import (
	"fmt"

	"github.com/jsdmitry/sushi-order/server/sql"
)

const connectionString = "root:12qwesdf@/sushi_order"

func main() {
	sqlProvider := sql.SQLDataProvider{ConnectionString: connectionString}
	for _, c := range sqlProvider.GetCategories() {
		fmt.Println(c)
	}
}
