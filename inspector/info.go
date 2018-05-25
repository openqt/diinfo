package inspector

import (
	"fmt"
	"github.com/heroku/docker-registry-client/registry"
	"log"
)

func InfoImage() []string {
	hub, err := registry.New(Settings.Registry, Settings.Username, Settings.Password)
	CheckError(err)

	repositories, err := hub.Repositories()
	CheckError(err)

	log.Print(repositories)
	return repositories
}

func ShowInfo() {
	for n, i := range ListImage() {
		fmt.Printf("%d.\t%s\n", n+1, i)
	}
}
