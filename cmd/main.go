package main

import (
	"github.com/andreashanson/golang-awake/internal/config"
	"github.com/andreashanson/golang-awake/internal/handlers"
	"github.com/andreashanson/golang-awake/internal/influencers"
	"github.com/andreashanson/golang-awake/internal/server"
	"github.com/andreashanson/golang-awake/pkg/postgres"
)

func main() {
	cfg := config.NewConfig()
	db, err := postgres.ConnectSQL(cfg)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	influencersRepo := postgres.NewInfluencersRepo(db)
	influencerSvc := influencers.NewService(influencersRepo)

	influencerHandlerl := handlers.NewInfluencerHandler(influencerSvc)

	newServer := server.NewServer("8080", influencerHandlerl)
	if err := newServer.Run(); err != nil {
		panic(err)
	}
}
