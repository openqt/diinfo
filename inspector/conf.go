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
	Wide int
	All  bool

	hub *registry.Registry
}

var Settings Setting

func CheckError(err error) {
	if err != nil {
		fmt.Println("ERROR:", err.Error())
		log.Fatal(err)
	}
}

func Configuration() {
	if !Settings.Verbose {
		log.SetFlags(0)
		log.SetOutput(ioutil.Discard)
	}

	if err := viper.ReadInConfig(); err != nil {
		log.Println(err)
	}
	log.Printf("Configuration: %s\n", viper.ConfigFileUsed())

	Settings.Registry = viper.GetString("registry")
	Settings.Username = viper.GetString("username")
	Settings.Password = viper.GetString("password")

	if Settings.Verbose {
		data, err := json.MarshalIndent(Settings, "", "  ")
		CheckError(err)
		fmt.Println(string(data))
	}

	hub, err := registry.NewInsecure(Settings.Registry, Settings.Username, Settings.Password)
	CheckError(err)
	Settings.hub = hub

	Settings.Json = viper.GetBool("json")
	Settings.Wide = viper.GetInt("wide")
	Settings.All = viper.GetBool("all")
}
