package inspector

import (
	"encoding/json"
	"fmt"
	"github.com/heroku/docker-registry-client/registry"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
)

type Setting struct {
	Verbose bool

	Registry string
	Username string
	Password string

	Json bool

	hub *registry.Registry
}

var Settings Setting

func CheckError(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

func Configuration() {
	Settings.Verbose = viper.GetBool("verbose")
	if !Settings.Verbose {
		log.SetFlags(0)
		log.SetOutput(ioutil.Discard)
	}

	log.Printf("Configuration: %s\n", viper.ConfigFileUsed())

	Settings.Registry = viper.GetString("registry")

	if Settings.Verbose {
		data, err := json.MarshalIndent(Settings, "", "  ")
		CheckError(err)
		fmt.Println(string(data))
	}

	hub, err := registry.New(Settings.Registry, Settings.Username, Settings.Password)
	CheckError(err)
	Settings.hub = hub
}
