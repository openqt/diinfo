package inspector

import (
    "fmt"
    "os"

    "github.com/spf13/cobra"
    "github.com/spf13/viper"
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = root()

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
    if err := rootCmd.Execute(); err != nil {
        fmt.Println(err)
        os.Exit(1)
    }
}

func init() {
    cobra.OnInitialize()
}

func root() *cobra.Command {
    viper.SetConfigName("diinfo")
    viper.AddConfigPath("/etc/diinfo")
    viper.AddConfigPath(".")

    rootCmd := cobra.Command{
        Use:   "diinfo",
        Short: "Docker Image Information Inspector",
        PersistentPreRun: func(cmd *cobra.Command, args []string) {
            Configuration()
        },
    }

    key := "verbose"
    rootCmd.PersistentFlags().BoolVarP(&Settings.Verbose, key, key[:1],
        false, "Show logs")
    viper.BindPFlag(key, rootCmd.PersistentFlags().Lookup(key))

    key = "registry"
    rootCmd.PersistentFlags().StringVarP(&Settings.Registry, key, key[:1],
        "", "Docker registry address")
    viper.BindPFlag(key, rootCmd.PersistentFlags().Lookup(key))

    key = "username"
    rootCmd.PersistentFlags().StringVarP(&Settings.Username, key, key[:1],
        "", "Password")
    viper.BindPFlag(key, rootCmd.PersistentFlags().Lookup(key))

    key = "password"
    rootCmd.PersistentFlags().StringVarP(&Settings.Password, key, key[:1],
        "", "Password")
    viper.BindPFlag(key, rootCmd.PersistentFlags().Lookup(key))

    return &rootCmd
}
