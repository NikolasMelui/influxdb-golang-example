package db

import (
	"context"
	"errors"
	influxdb2 "github.com/influxdata/influxdb-client-go/v2"
	"os"
)

func ConnectToDB() (influxdb2.Client, error) {
	dbToken := os.Getenv("INFLUXDB_TOKEN")
	if dbToken == "" {
		return nil, errors.New("INFLUXDB_TOKEN must be set")
	}

	dbURL := os.Getenv("INFLUXDB_URL")
	if dbURL == "" {
		return nil, errors.New("INFLUXDB_URL must be set")
	}

	client := influxdb2.NewClient(dbURL, dbToken)

	_, err := client.Health(context.Background())
	return client, err
}
