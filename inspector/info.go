package inspector

import (
	"encoding/json"
	"fmt"
	"github.com/olekukonko/tablewriter"
	"io/ioutil"
	"os"
	"strings"
	"bytes"
)

type Config struct {
	Tty        bool
	Env        []string
	Cmd        []string
	Image      string
	Volume     map[string]interface{}
	WorkingDir string
	Entrypoint []string
	Labels     []string
}

type History struct {
	Created   string
	CreatedBy string `json:"created_by"`
	Empty     bool   `json:"empty_layer"`

	Digest  string
	Size    int64
	Command string
}

type FullConfig struct {
	Architecture  string
	Config        Config
	Created       string
	DockerVersion string `json:"docker_version"`
	History       []History
	Os            string
}

func InfoImage(repo, tag string) FullConfig {
	manifest, err := Settings.hub.ManifestV2(repo, tag)
	CheckError(err)

	reader, err := Settings.hub.DownloadLayer(repo, manifest.Config.Digest)
	CheckError(err)

	data, err := ioutil.ReadAll(reader)
	CheckError(err)

	config := FullConfig{}
	json.Unmarshal(data, &config)

	t, _ := json.MarshalIndent(config, "", "  ")
	fmt.Println(string(t))

	return config
}

func InfoImageOrigin(repo, tag string) {
	manifest, err := Settings.hub.ManifestV2(repo, tag)
	CheckError(err)

	data, err := manifest.MarshalJSON()
	CheckError(err)

	fmt.Println("---------- Manifest ----------")
	fmt.Println(string(data))

	reader, err := Settings.hub.DownloadLayer(repo, manifest.Config.Digest)
	CheckError(err)

	data, err = ioutil.ReadAll(reader)
	CheckError(err)

	var buf bytes.Buffer
	err = json.Indent(&buf, data, "", "  ")
	CheckError(err)

	fmt.Println("---------- Configuration ----------")
	fmt.Println(buf.String())
}

func ShowInfo(args []string) {
	for _, arg := range args {
		var repo, tag string

		ns := strings.SplitN(arg, ":", 1)
		repo = ns[0]
		if len(ns) > 1 {
			tag = ns[1]
		} else {
			tag = "latest"
		}

		if Settings.Json {
			InfoImageOrigin(repo, tag)
		} else {
			InfoImage(repo, tag)

			table := tablewriter.NewWriter(os.Stdout)
			table.SetHeader([]string{"No", "Size", "Command"})
			table.SetFooter([]string{"Total", "", ""})
			table.Render()
		}

		//for n, i := range InfoImage(repo, tag) {
		//	table.Append(i)
		//	fmt.Printf("%d.\t%s\n", n+1, i)
		//}
	}
}
