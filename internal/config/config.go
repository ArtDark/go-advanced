package config

import (
	"log"
	"os"

	"github.com/joho/godotenv"
)


const dbDsn = "DB_DSN"
const secret = "TOKEN"

type Main struct {
	Db Db
	Auth Auth
}

type Db struct {
	Dsn string
}

type Auth struct {
	Secret string
}

func New() *Main {
	return &Main{
		Db: Db{
		},
		Auth: Auth{
			
		},
	}
}

func (m *Main) Load() *Main {
	if err := godotenv.Load(".env"); err != nil {
		log.Printf("error: %v", err)
	}
	return m
}

func (m *Main) Init() *Main {
	m.Db.Dsn = os.Getenv(dbDsn)
	m.Auth.Secret = os.Getenv(secret)
	return m
}