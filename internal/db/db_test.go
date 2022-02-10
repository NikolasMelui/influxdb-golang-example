package db

import (
	"testing"
)

func Test_GetConnection(t *testing.T) {
	connection := GetConnection()
	if connection != "__connection__" {
		t.Errorf("GetConnection() error. Something went wrong.")
		return
	}
}
