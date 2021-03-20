package database

import (
	"testing"
)

func TestGetLastRecord(t *testing.T) {
	InsertNewRecord(3, "test value")
	GetLastRecord()
	DeleteLastRecord(3)
	GetLastRecord()
}
