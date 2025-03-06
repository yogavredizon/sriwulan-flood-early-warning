package utils_test

import (
	"fmt"
	"log"
	"testing"

	"github.com/yogavredizon/sriwulan-flood-early-warning/utils"
)

func TestFetchYAML(t *testing.T) {
	dsn, err := utils.FetchYAML()
	if err != nil {
		log.Panic(err)
	}

	fmt.Println(dsn)
}
