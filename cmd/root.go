package cmd

import (
	"github.com/spf13/cobra"
	"leizhenpeng/go-gpt3-cli/services"
	"os"
)

const (
	keyName = "chat-go-key"
	Version = "0.0.1"
	AppName = "chat-go"
)

var keyMsg *services.KeyMag

var rootCmd = &cobra.Command{
	Use:   AppName,
	Short: "CLI ChatBot Power By Gpt3",
	Run: func(cmd *cobra.Command, args []string) {
		if flag := cmd.Flag("version"); flag != nil && flag.Value.String() == "true" {
			cmd.Println(AppName, "v"+Version)
		} else {
			cmd.Help()
		}
	},
}

func Execute() {
	err := rootCmd.Execute()

	if err != nil {
		os.Exit(1)
	}
}

func init() {
	services.NewKeyMag()
	rootCmd.Flags().BoolP("version", "v", false, "version of "+AppName)
}
