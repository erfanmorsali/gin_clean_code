package main

import (
	"clean_code/api"
	"clean_code/database"
)

func main() {
	db, err := database.ConnectToDataBase()
	if err != nil {
		panic(err)
	}
	if err := api.RunApiServer(db); err != nil {
		panic(err)
	}
}
