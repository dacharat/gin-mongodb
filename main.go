package main

import (
	"github.com/dacharat/gin-mongodb/database"
	"github.com/dacharat/gin-mongodb/routers"
)

func main() {
	database.Init()

	router := routers.GenerateRouter()
	router.Run()
}
