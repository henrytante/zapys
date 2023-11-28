package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)

func LoadDotEnv() error {
	err := godotenv.Load()
	if err != nil{
		log.Fatalf("Erro ao carregar dotenv. Erro: %s", err)
		return err
	}
	return nil
}
func GetVarEnv(key string) string {
	return os.Getenv(key)
}