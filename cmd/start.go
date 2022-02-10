package main

import (
	"fmt"
	"github.com/NikolasMelui/influxdb-golang-example/internal/db"
)

func init() {
	fmt.Println("Initializing the project...")
}

func main() {
	connection := db.GetConnection()
	fmt.Printf("Connection '%s' was successfully get!", connection)
}
