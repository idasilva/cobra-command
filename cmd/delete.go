package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var deleteCmd = &cobra.Command{
	Use:   "delete",
	Short: "A brief description of your command",
	Long:  "A longer description",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("Delete command")

	},
}

func init()  {
	rootCmd.AddCommand(deleteCmd)
	deleteCmd.Flags().StringP("delete", "d", "", "delete something")
	
}
