/*
Copyright © 2024 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"log"

	"github.com/notnmeyer/dngl/internal/httpclient"

	"github.com/spf13/cobra"
)

// getCmd represents the get command
var getCmd = &cobra.Command{
	Use:   "get",
	Short: "get a note by id",
	Long: `get a note by id:

dngl get <id>`,
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]
		url := fmt.Sprintf("/%s/%s", "note", id)
		resp, err := httpclient.NewRequest("GET", url, nil)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(resp))
	},
}

func init() {
	rootCmd.AddCommand(getCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// getCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// getCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
