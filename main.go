package main

import (
	"fmt"
	"log"
	"os"

	api "github.com/ntwaliheritier/go_blog_app/api"
	storage "github.com/ntwaliheritier/go_blog_app/storage"
)

func main() {

	config := storage.Config {
		Host: os.Getenv("Host"),
		Port: os.Getenv("Port"),
		Username: os.Getenv("Username"),
		Password: os.Getenv("Password"),
		DBName: os.Getenv("DBName"),
		SSLMode: os.Getenv("SSLMode"),
	}

	if err := storage.NewConnection(config); err != nil {
		fmt.Println(err)
		log.Fatal("Invalid database credentials")
	}

	app := api.App()
	log.Fatal(app.Listen(":3000"))
}