package utils_test

import (
	"fmt"
	"log"
	"testing"

	"github.com/yogavredizon/sriwulan-flood-early-warning/utils"
)

func TestEncDec(t *testing.T) {
	iv := "my16digitIvKey12"
	key := "my32digitkey12345678901234567890"
	plainText := "root"

	enc, err := utils.Encrypt([]byte(plainText), []byte(key), []byte(iv))
	if err != nil {
		log.Panic(err)
	}

	fmt.Println(enc)
	dec, err := utils.Decrypt(enc, []byte(key), []byte(iv))
	if err != nil {
		log.Panic(err)
	}
	fmt.Println("Dec", dec)
}
