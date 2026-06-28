package cmd

import (
	"fmt"
	"os"
	"os/user"

	"github.com/ccaroon/mnky/repl"
	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:     "mnky",
	Short:   "A Monkey Interpreter",
	Version: version, // See version.go
	// PersistentPreRun: func(cmd *cobra.Command, args []string) {
	// 	if targetFlag != "" {
	// 		if config.TargetExists(targetFlag) {
	// 			config.SetTarget(targetFlag)
	// 		} else {
	// 			handleCmdError(
	// 				fmt.Errorf("Target [%s] does not exist!", targetFlag),
	// 			)
	// 		}
	// 	}
	// },
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) == 0 {
			user, err := user.Current()
			if err != nil {
				panic(err)
			}

			fmt.Printf("Monkey v%s | Greetings %s\n", version, user.Username)
			fmt.Printf("Ctrl-D to exit\n")
			repl.Start(os.Stdin, os.Stdout)
		} else {
			fmt.Println("Interpreter NOT implemented yet!")
		}
	},
}

func init() {
	rootCmd.SetVersionTemplate("{{.DisplayName }} v{{ .Version }}\n")

	// cobra.OnInitialize(config.InitConfig)

	rootCmd.AddCommand(
		testCmd,
		versionCmd,
	)
}

func handleCmdError(err error) {
	if err != nil {
		fmt.Fprintln(os.Stderr, err.Error())
		os.Exit(1)
	}
}

func Execute() {
	err := rootCmd.Execute()
	handleCmdError(err)
}
