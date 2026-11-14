package main

import (
	"log"

	"github.com/maiconDeSouza/lotofacil_lab_v3/internal/database"
	"github.com/maiconDeSouza/lotofacil_lab_v3/internal/env"
	"github.com/maiconDeSouza/lotofacil_lab_v3/internal/handler"
	"github.com/maiconDeSouza/lotofacil_lab_v3/internal/lottery"
	"github.com/maiconDeSouza/lotofacil_lab_v3/internal/repo"
	"github.com/maiconDeSouza/lotofacil_lab_v3/internal/server"
	"github.com/maiconDeSouza/lotofacil_lab_v3/internal/service"
)

func main() {
	env, errAppRes := env.InitENV()
	if errAppRes != nil {
		log.Fatal(errAppRes)
	}

	lottery := lottery.NewLottery(env)
	collection, errAppRes := database.NewConnectMongo(env, lottery)
	if errAppRes != nil {
		log.Fatal(errAppRes)
	}

	repoGame := repo.NewGameRepo(collection.CollectionGamers)
	repoFixed := repo.NewFixedRepo(collection.CollectionFixed)
	repoScored := repo.NewScoredRepo(collection.CollectionsScored)

	service := service.NewGameService(repoGame, repoFixed, repoScored, lottery)
	handler := handler.NewHandler(service)
	server := server.NewServer(handler, env)

	errAppRes = server.Run()
	if errAppRes != nil {
		log.Fatal(errAppRes)
	}
}
