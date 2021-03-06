package cmd

import (
	"fmt"

	"github.com/ericbaukhages/choose/choose"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/cobra"
)

var (
	openNewSession = false
)

// newCmd represents the new command
var newCmd = &cobra.Command{
	Use:     "new",
	Short:   "Creates a new tmux session",
	Long:    ``,
	Aliases: []string{"add"},
	Run: func(cmd *cobra.Command, args []string) {
		configFileName, err := homedir.Expand("~/.tmux.sessions.log")
		if err != nil {
			fmt.Printf("Prompt failed: %v\n", err)
		}

		config := choose.Config{
			Location: configFileName,
		}
		config.Parse()

		var (
			name string
			path string
		)

		if len(args) == 2 {
			name = args[0]
			path = args[1]
		} else {
			cmd.Help()
			return
		}

		err = config.Add(name, path)
		if err != nil {
			fmt.Printf("New session could not be added: %v\n", err)
			return
		}

		err = config.Save()
		if err != nil {
			fmt.Printf("Project could not be saved: %v\n", err)
			return
		}

		if openNewSession {
			session := choose.Session{
				Path:    config.Values[name],
				Session: name,
			}
			_, err = session.Start()
			if err != nil {
				fmt.Printf("Sesssion failed %v\n", err)
				return
			}
		}
	},
}

func init() {
	rootCmd.AddCommand(newCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// newCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// newCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
	newCmd.Flags().BoolVarP(&openNewSession, "open", "o", true, "Open new session after creation")
}
