package inspector

import (
    "fmt"

    "github.com/spf13/cobra"
    "runtime"
)

var (
    AppVersion string
    GoVersion  string
    GitVersion string
    BuildTime  string
)

func init() {
    // versionCmd represents the appVersion command
    var versionCmd = &cobra.Command{
        Use:   "version",
        Short: "Show Version",
        Long:  `Print the appVersion information.`,
        Run: func(cmd *cobra.Command, args []string) {
            fmt.Printf("ReadKeeper:     %s\n", AppVersion)
            fmt.Printf("GitVersion:     %s\n", GitVersion)
            fmt.Printf("GoCompiler:     %s\n", GoVersion)
            fmt.Printf("Build Time:     %s\n", BuildTime)
            fmt.Printf("Go Version:     %s\n", runtime.Version())
        },
    }

    rootCmd.AddCommand(versionCmd)
}
