package cmd

// Context holds the information from the previous query

import (
	"fmt"
	"leizhenpeng/go-gpt3-cli/services"

	"github.com/spf13/cobra"
)

type Context struct {
	PreviousQuery string
}

// runCmd represents the run command
var runCmd = &cobra.Command{
	Use:   "run",
	Short: "Run a gp3 chat bot",
	Run: func(cmd *cobra.Command, args []string) {
		keyMsg = services.GetKeyMag()
		key := keyMsg.GetKey(keyName)
		if key == "" {
			fmt.Println(`Don't find open-api key, please set your key first.
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

	rootCmd.AddCommand(runCmd)

	runCmd.Flags().BoolP("interactive", "i", true, "interactive mode")
	runCmd.Flags().StringP("prompt", "p", "hello gopher", "prompt mode")
}

func InteractiveMode() {
	//services.SetContext(ctx)
	history := services.NewCacheHistory()
	fmt.Print("Welcome to the GPT-3 chat bot. \nType 'exit' to quit, Type 'clear' to clear the context. \n")
	for {
		question := services.AskUserQuestion()
		if question == "exit" {
			break
		}
		if question == "clear" {
			history.ClearQACache()
			continue
		}
		cache, b := history.GetQACache()
		if b {
			question = cache + "/n" + question
		}
		reply, ok := services.GetAnswer(question)

		if ok {
			history.SetQACache(question, reply)

		}
	}
}
