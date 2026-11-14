package env

import (
	"os"

	"github.com/joho/godotenv"
	"github.com/maiconDeSouza/lotofacil_lab_v3/internal/errapp"
)

type Env struct {
	SERVER_PORT string

	MONGO_DB_NAME           string
	MONGO_COLLECTION_FIXED  string
	MONGO_COLLECTION_GAMERS string
	MONGO_COLLECTION_SCORED string
	MONGO_ROOT_USER         string
	MONGO_ROOT_PASSWORD     string
	MONGO_PORT              string

	URL string
}

func InitENV() (*Env, *errapp.ErrAppRes) {
	if err := godotenv.Load(); err != nil {
		errAppRes := errapp.NewErrAppRes(err, "❌ Erro ao carregar arquivo .env", 0)
		return nil, errAppRes
	}

	e := Env{
		SERVER_PORT:             os.Getenv("SERVER_PORT"),
		MONGO_DB_NAME:           os.Getenv("MONGO_DB_NAME"),
		MONGO_COLLECTION_FIXED:  os.Getenv("MONGO_COLLECTION_FIXED"),
		MONGO_COLLECTION_GAMERS: os.Getenv("MONGO_COLLECTION_GAMERS"),
		MONGO_COLLECTION_SCORED: os.Getenv("MONGO_COLLECTION_SCORED"),
		MONGO_ROOT_USER:         os.Getenv("MONGO_ROOT_USER"),
		MONGO_ROOT_PASSWORD:     os.Getenv("MONGO_ROOT_PASSWORD"),
		MONGO_PORT:              os.Getenv("MONGO_PORT"),
		URL:                     os.Getenv("URL"),
	}

	return &e, nil
}
