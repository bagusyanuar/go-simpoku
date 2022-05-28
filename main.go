package main

import (
	"flag"
	"fmt"
	"go-simpoku/database"
	"go-simpoku/routes"
	"os"

	"github.com/joho/godotenv"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)


func main() {
	err := godotenv.Load()
	if err != nil {
		panic("Error : Failed Load .ENV File")
	}
	appPort := os.Getenv("APP_PORT")
	database.DB, err = gorm.Open(mysql.Open(database.Build()))
	if err != nil {
		panic("Error : Failed To Connect Database")
	}
	fmt.Println("success connect to database")
	migrate := flag.String("m", "", "Unsupport Command")
	flag.Parse()
	command := *migrate

	if command == "migrate" {
		database.Migrate()
		fmt.Printf("Successfull Migrate")
		return
	}
	server := routes.InitRoutes()
	port := fmt.Sprintf(":%s", appPort)
	server.Run(port)
}