package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

func init() {

	//cobra.OnInitialize(initConfig)
	//RootCmd.PersistentFlags().StringVar(&cfgFile, "config", "c", "", "config file (default is ./funky.yaml)")
	//RootCmd.PersistentFlags().StringVarP(&projectBase, "projectbase", "b", "", "base project directory eg. github.com/spf13/")
	//RootCmd.PersistentFlags().StringP("author", "a", "YOUR NAME", "Author name for copyright attribution")
	//RootCmd.PersistentFlags().StringVarP(&userLicense, "license", "l", "", "Name of license for the project (can provide `licensetext` in config)")
	//RootCmd.PersistentFlags().Bool("viper", true, "Use Viper for configuration")
	//viper.BindPFlag("author", RootCmd.PersistentFlags().Lookup("author"))
	//viper.BindPFlag("projectbase", RootCmd.PersistentFlags().Lookup("projectbase"))
	//viper.BindPFlag("useViper", RootCmd.PersistentFlags().Lookup("viper"))
	//viper.SetDefault("author", "NAME HERE <EMAIL ADDRESS>")
	//viper.SetDefault("license", "apache")
}

// RootCmd is the top-level cli command handler
var RootCmd = &cobra.Command{
	Use:   "funky",
	Short: "Funky manages Google Cloud Functions in multiple languages",
	Long: `A simple and intuitive manager for Google Cloud Functions
that can be written in different languages.`,
	Run: func(cmd *cobra.Command, args []string) {
		// read the 'funky.yaml' config file
		viper.SetConfigName("funky")
		viper.AddConfigPath(".")
		if err := viper.ReadInConfig(); err != nil {
			fmt.Printf("\nNo 'funky.yaml' config file found\n\n")
		}
	},
}
