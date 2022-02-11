package main

import (
	"fmt"
	"github.com/NikolasMelui/influxdb-golang-example/internal/db"
	"github.com/joho/godotenv"
)

func init() {
	fmt.Println("Initializing the project...")
	err := godotenv.Load("build/influxdb.env")
	if err != nil {
		panic(fmt.Errorf("cannot read the influxdb.env file"))
	}
}

func main() {
	connection, err := db.ConnectToDB()
	if err != nil {
		panic(err)
	}
	fmt.Printf("Successfully connected to DB:\n")
	fmt.Printf("%s", connection)
}
