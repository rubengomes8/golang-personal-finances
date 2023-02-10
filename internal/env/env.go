package env

import "github.com/joho/godotenv"

func init() {
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
}
