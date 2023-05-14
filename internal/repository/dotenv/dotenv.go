package dotenv

import (
	"log"
	"os"
	"strings"

	"github.com/joho/godotenv"
)

func Load() {

	curDir, err := os.Getwd()
	if err != nil {
		log.Println(err)
	}

	if strings.Contains(curDir, "cmd") == true {
		curDir = curDir + "/.."
	}

	err = godotenv.Load(curDir + "/.env")
	if err != nil {
		log.Fatalf("Error loading .env file")
	}
}

func Get(key string) string {
	return os.Getenv(key)
}
