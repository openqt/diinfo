package main

import (
	"github.com/openqt/diinfo/inspector"
	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"log"
)

func root() *cobra.Command {
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

	key = "registry"
	rootCmd.PersistentFlags().StringVarP(&inspector.Settings.Registry, key, key[:1],
		"", "Docker registry address")
	viper.BindPFlag(key, rootCmd.PersistentFlags().Lookup(key))

	return &rootCmd
}

func ls() *cobra.Command {
	listCmd := cobra.Command{
		Use:   "ls",
		Short: "List images in docker registry",
		Run: func(cmd *cobra.Command, args []string) {
			inspector.ShowList()
		},
	}
	return &listCmd
}

func info() *cobra.Command {
	infoCmd := cobra.Command{
		Use:   "info",
		Short: "Show docker image internals",
		Args:  cobra.MinimumNArgs(1),
		Run: func(cmd *cobra.Command, args []string) {
			inspector.ShowInfo(args)
		},
	}

	key := "json"
	infoCmd.PersistentFlags().BoolVarP(&inspector.Settings.Json, key, key[:1],
		false, "Show data in JSON style")
	viper.BindPFlag(key, infoCmd.PersistentFlags().Lookup(key))

	return &infoCmd
}

func main() {
	rootCmd := root()
	rootCmd.AddCommand(
		ls(),
		info(),
	)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
