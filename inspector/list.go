package inspector

import (
	"fmt"
	"log"
)

func ListImage() []string {
	repositories, err := Settings.hub.Repositories()
	CheckError(err)

	log.Print(repositories)
	return repositories
}

func ListTag(repo string) []string {
	tags, err := Settings.hub.Tags(repo)
	CheckError(err)

	log.Print(tags)
	return tags
}

func ShowList() {
	for n, i := range ListImage() {
		tags := ListTag(i)
		fmt.Printf("%d.\t%s: %v\n", n+1, i, tags)
	}
}
