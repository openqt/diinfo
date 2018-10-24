package inspector

import (
    "fmt"
    "log"
    "github.com/spf13/cobra"
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

func init() {
    listCmd := &cobra.Command{
        Use:   "ls",
        Short: "List images in docker registry",
        Run: func(cmd *cobra.Command, args []string) {
            ShowList()
        },
    }
    rootCmd.AddCommand(listCmd)
}
