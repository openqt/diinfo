package main

import (
	"github.com/openqt/diinfo/inspector"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
)

func main() {
	viper.SetConfigName("diinfo")
	viper.AddConfigPath("/etc/diinfo")
	viper.AddConfigPath(".")
	inspector.CheckError(viper.ReadInConfig())

	rootCmd := cobra.Command{
		Use:   "diinfo",
		Short: "Docker Image Information Inspector",
		PersistentPreRun: func(cmd *cobra.Command, args []string) {
			inspector.Configuration()
		},
	}
	key := "verbose"
	rootCmd.PersistentFlags().BoolVarP(&inspector.Settings.Verbose, key, key[:1],
		false, "Show logs")
	viper.BindPFlag(key, rootCmd.PersistentFlags().Lookup(key))

	listCmd := cobra.Command{
		Use:   "ls",
		Short: "List images in docker registry",
		Run: func(cmd *cobra.Command, args []string) {
			inspector.ShowList()
		},
	}

	key = "registry"
	listCmd.PersistentFlags().StringVarP(&inspector.Settings.Registry, key, key[:1],
		"", "Docker registry address")
	viper.BindPFlag("Registry", listCmd.PersistentFlags().Lookup(key))
	rootCmd.AddCommand(
		&listCmd,
	)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
