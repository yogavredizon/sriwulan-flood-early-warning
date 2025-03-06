package db_test

import (
	"fmt"
	"log"
	"testing"

	"github.com/yogavredizon/sriwulan-flood-early-warning/db"
	"github.com/yogavredizon/sriwulan-flood-early-warning/utils"
)

func TestConn(t *testing.T) {
	dsn, err := utils.FetchYAML()
	if err != nil {
		log.Fatal(err)
	}

	_, err = db.Conn(dsn)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected!")
}
