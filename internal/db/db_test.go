package db

import (
	"context"
	"github.com/influxdata/influxdb-client-go/v2/domain"
	"github.com/joho/godotenv"
	"testing"
)

func Test_ConnectToDB(t *testing.T) {
	err := godotenv.Load("../../build/influxdb.env")
	if err != nil {
		t.Errorf("Cannot read the .env file")
		return
	}

	tests := []struct {
		name    string
		wantErr bool
	}{
		{
			name:    "Successful connection to InfluxDB",
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := ConnectToDB()
			if (err != nil) != tt.wantErr {
				t.Errorf("ConnectToInfluxDB() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			health, err := got.Health(context.Background())
			if (err != nil) && health.Status == domain.HealthCheckStatusPass {
				t.Errorf("connectToInfluxDB() error. database not healthy")
				return
			}
			got.Close()
		})
	}
}
