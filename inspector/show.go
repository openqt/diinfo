package inspector

import (
    "bytes"
    "encoding/json"
    "fmt"
    "github.com/opencontainers/go-digest"
    "io/ioutil"
    "regexp"
    "strings"
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

    Digest string
    Size   int64
}

type FullConfig struct {
    Architecture  string
    Config        Config
    Created       string
    DockerVersion string `json:"docker_version"`
    History       []History
    Os            string
}

func getConfig(repo string, digest digest.Digest) []byte {
    reader, err := Settings.hub.DownloadLayer(repo, digest)
    CheckError(err)
    data, err := ioutil.ReadAll(reader)
    CheckError(err)

    return data
}

func InfoImage(repo, tag string) FullConfig {
    manifest, err := Settings.hub.ManifestV2(repo, tag)
    CheckError(err)

    config := FullConfig{}
    json.Unmarshal(getConfig(repo, manifest.Config.Digest), &config)

    for i, j := 0, 0; i < len(config.History); i++ {
        history := &config.History[i]
        if !history.Empty {
            layer := manifest.Layers[j]
            history.Digest = string(layer.Digest)
            history.Size = layer.Size

            j++
        }
    }
    return config
}

func InfoImageOrigin(repo, tag string) {
    manifest, err := Settings.hub.ManifestV2(repo, tag)
    CheckError(err)

    data, err := manifest.MarshalJSON()
    CheckError(err)

    fmt.Println("---------- Manifest ----------")
    fmt.Println(string(data))

    var buf bytes.Buffer
    err = json.Indent(&buf, getConfig(repo, manifest.Config.Digest), "", "  ")
    CheckError(err)

    fmt.Println("---------- Configuration ----------")
    fmt.Println(buf.String())
}

func splitLastColon(s string) (string, string) {
    i := strings.LastIndex(s, ":")
    if i > 0 {
        return s[:i], s[i+1:]
    }
    return s, "latest" // default tag
}

func ShowInfo(args []string) {
    for _, arg := range args {
        repo, tag := splitLastColon(arg)

        if Settings.Json {
            InfoImageOrigin(repo, tag)
        } else {
            var total int64

            InfoImage(repo, tag)

            config := InfoImage(repo, tag)

            // print headers
            fmt.Printf("%v.\t% 10v\t%s\n", "No", "Size", "Command [/bin/sh -c]")
            fmt.Printf("%v\t% 10v\t%s\n", "---", "----", "--------------------")
            n := 0
            for _, i := range config.History {
                cmd := i.CreatedBy
                if Settings.Wide > 0 { // show with limited length
                    reg := regexp.MustCompile(`\s+`) // normalize control characters
                    cmd = reg.ReplaceAllString(cmd, " ")

                    if strings.HasPrefix(cmd, "/bin/sh -c ") {
                        cmd = cmd[11:] // remove fix prefix to reduce length
                    }
                    if len(cmd) > Settings.Wide { // truncate and mark with ...
                        cmd = cmd[:Settings.Wide] + " ..."
                    }
                }
                if Settings.All || i.Size > 0 { // show non-zero layers by default
                    total += i.Size
                    fmt.Printf("%v.\t% 10v\t%s\n", n+1, i.Size, cmd)
                    n++
                }
            }
            // print total size by MB
            fmt.Printf("------------------\n")
            fmt.Printf("%s image size: %.2f MB\n\n", arg, float64(total)/1024/1024)
        }
    }
}
