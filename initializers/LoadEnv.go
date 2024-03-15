package initializers

import (
    "github.com/joho/godotenv"
    "log"
)

func LoadEnv() {
    // Load .env file
    if err := godotenv.Load(); err != nil {
        log.Fatalf("Error loading .env file: %v", err)
    }
}
