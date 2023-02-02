package cmd

import (
	"fmt"
	"leizhenpeng/go-gpt3-cli/services"

	"github.com/spf13/cobra"
)

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a gp3 chat bot",
	Run: func(cmd *cobra.Command, args []string) {
		keyMsg = services.GetKeyMag()
		key := keyMsg.GetKey(keyName)
		if key == "" {
			fmt.Println(`You don't have a key, please set your key first.
Find your key from https://beta.openai.com/account/api-keys.
Run command : go-chat key -s <your key>`)
			return
		}
		services.InitClient(key)
		if cmd.Flag("interactive").Value.String() == "true" {
			InteractiveMode()
		} else if cmd.Flag("prompt").Value.String() != "" {
			services.GetAnswer(cmd.Flag("prompt").Value.String())
		} else {
			cmd.Help()
		}

	},
}

func init() {
	//check if the key is exist

	//if not, ask user to input the key

	rootCmd.AddCommand(runCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	//runCmd.PersistentFlags().StringP("foo", "f", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	runCmd.Flags().BoolP("interactive", "i", true, "interactive mode")
	runCmd.Flags().StringP("prompt", "p", "hello gopher", "prompt mode")
}

func InteractiveMode() {
	fmt.Print("Welcome to the GPT-3 chat bot. Type 'exit' to quit.\n")
	for {
		question := services.AskUserQuestion()
		if question == "exit" {
			break
		}
		services.GetAnswer(question)
	}
}
