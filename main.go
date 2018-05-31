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
		Use:   "show",
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

	key = "wide"
	infoCmd.PersistentFlags().IntVarP(&inspector.Settings.Wide, key, key[:1],
		0, "Length of command line (0 is no limited)")
	viper.BindPFlag(key, infoCmd.PersistentFlags().Lookup(key))

	key = "all"
	infoCmd.PersistentFlags().BoolVarP(&inspector.Settings.All, key, key[:1],
		false, "Show zero size layers")
	viper.BindPFlag(key, infoCmd.PersistentFlags().Lookup(key))

	return &infoCmd
}

func del() *cobra.Command {
	delCmd := cobra.Command{
		Use:   "del",
		Short: "Delete image(s)",
		Run: func(cmd *cobra.Command, args []string) {
			inspector.DelImages(args)
		},
	}
	return &delCmd
}

func main() {
	rootCmd := root()
	rootCmd.AddCommand(
		ls(),
		info(),
		del(),
	)

	if err := rootCmd.Execute(); err != nil {
		log.Fatal(err)
	}
}
