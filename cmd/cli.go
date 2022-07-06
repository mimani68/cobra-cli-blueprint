package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var status *string
var username *string
var param *string
var password *string

var rootCmd = &cobra.Command{}

var restApiCmd = &cobra.Command{
	Use:     "api",
	Aliases: []string{"rest", "r", "rest-api"},
	Short:   "fast exposing api",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Running API application")
		fmt.Println(*status)
	},
}

// $ a console --username mahdi --password 123
var consoleCmd = &cobra.Command{
	Use:     "console",
	Aliases: []string{"c", "cli"},
	Short:   "Console envirnoment",
	Args: func(cmd *cobra.Command, args []string) error {
		// $ a console --username mahdi --password 123 db-list
		if len(args) > 0 && args[0] == "db-list" {
			param = new(string)
			*param = "(1) mysql"
		} else {
			param = new(string)
			*param = "[x]"
		}
		return nil
	},
	PreRun: func(cmd *cobra.Command, args []string) {
		fmt.Println("Run console application")
	},
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Welcome %s", *username)
		fmt.Printf("\n%s", *param)
	},
}

// Child command `console` > `db`
// a console --username mahdi --password 123 db
var consoleDbCmd = &cobra.Command{
	Use:   "db",
	Short: "Managing database",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Run database managment application")
	},
}

func init() {
	status = new(string)
	restApiCmd.PersistentFlags().StringVar(status, "status", "UNKNOWN", "Check status of server")

	username = new(string)
	password = new(string)
	consoleCmd.PersistentFlags().StringVar(username, "username", "", "User identity username")
	consoleCmd.PersistentFlags().StringVar(password, "password", "", "User cerdential")
	consoleCmd.MarkPersistentFlagRequired("username")
	consoleCmd.MarkPersistentFlagRequired("password")

	rootCmd.AddCommand(restApiCmd)
	rootCmd.AddCommand(consoleCmd)

	consoleCmd.AddCommand(consoleDbCmd)
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}
