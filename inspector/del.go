package inspector

import (
	"fmt"
	"strings"
)

func DelImage(repo, tag string) {
	manifest, err := Settings.hub.ManifestV2(repo, tag)
	CheckError(err)

	err = Settings.hub.DeleteManifest(repo, manifest.Config.Digest)
	CheckError(err)

	fmt.Printf("%v:%v deleted.", repo, tag)
}

func DelImages(args []string) {
	for _, arg := range args {
		var repo, tag string

		ns := strings.SplitN(arg, ":", 1)
		repo = ns[0]
		if len(ns) > 1 {
			tag = ns[1]
		} else {
			tag = "latest" // default tag
		}

		DelImage(repo, tag)
	}
}
