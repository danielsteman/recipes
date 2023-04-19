package main

import (
	"fmt"
	"os"

	"github.com/joho/godotenv"
)

func main() {
	godotenv.Load()
	key := os.Getenv("API_KEY")
	fmt.Println(key)
}
