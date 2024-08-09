/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"
	"strings"

	"github.com/notnmeyer/dngl/internal/httpclient"

	"github.com/spf13/cobra"
)

// createCmd represents the create command
var createCmd = &cobra.Command{
	Use:   "create",
	Short: "Create a note",
	Long: `Cretae a new note:

dngl create "something interesting"`,
	Run: func(cmd *cobra.Command, args []string) {
		// mindlessly turn all the args into a string
		content := strings.Join(args, " ")
		resp, err := httpclient.NewRequest("POST", "/note/create", &content)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(resp))
	},
}

func init() {
	rootCmd.AddCommand(createCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// createCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// createCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
