package inspector

import (
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"io/ioutil"
	"log"
)

type Setting struct {
	Verbose bool

	Registry string
	Username string
	Password string
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
}
