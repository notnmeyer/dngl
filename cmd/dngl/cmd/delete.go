package cmd

import (
	"fmt"
	"log"

	"github.com/notnmeyer/dngl/internal/httpclient"

	"github.com/spf13/cobra"
)

// deleteCmd represents the delete command
var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "delete a note",
	Long: `delete a note

dngl delete <id>`,
	Run: func(cmd *cobra.Command, args []string) {
		id := args[0]
		resp, err := httpclient.NewRequest("POST", "/note/delete/"+id, nil)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Println(string(resp))
	},
}

func init() {
	rootCmd.AddCommand(deleteCmd)

	// Here you will define your flags and configuration settings.

	// Cobra supports Persistent Flags which will work for this command
	// and all subcommands, e.g.:
	// deleteCmd.PersistentFlags().String("foo", "", "A help for foo")

	// Cobra supports local flags which will only run when this command
	// is called directly, e.g.:
	// deleteCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}
