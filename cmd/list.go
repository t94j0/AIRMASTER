package cmd

import (
	"fmt"

	"github.com/t94j0/AIRMASTER/domain"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
)

var purchase bool
var file string
var keywords []string

// listCmd represents the list command
var listCmd = &cobra.Command{
	Use:   "list",
	Short: "List domains to purchase",
	Long:  `List domains and have the option to purchase the domains as well`,
	Run: func(cmd *cobra.Command, args []string) {
		// Configure Namecheap client
		/*
			apiUser, apiToken, username := viper.GetString("apiUser"), viper.GetString("apiToken"), viper.GetString("username")
			if apiUser == "" || apiToken == "" || username == "" {
				fmt.Println("Please specify namecheap API tokens")
				os.Exit(1)
			}

			ncClient := namecheap.NewClient(apiUser, apiToken, username)
		*/

		// Configure domain finding mechanism
		if viper.GetString("file") != "" {
			if err := domain.ParseFile(viper.GetString("file")); err != nil {
				fmt.Println(err)
			}
		} else if len(viper.GetStringSlice("keyword")) != 0 {
			if err := domain.ParseKeywords(viper.GetStringSlice("keyword")); err != nil {
				fmt.Println(err)
			}
		} else {
			fmt.Println("Please specify either a file or keywords")
			fmt.Println(cmd.Usage())
		}

	},
}

func init() {
	RootCmd.AddCommand(listCmd)

	listCmd.Flags().BoolP("purchase", "p", false, "Purchase domains that are listed")
	listCmd.Flags().StringP("file", "f", "", "File used for checking domains")
	listCmd.Flags().StringSliceP("keyword", "k", nil, "Keyword for searching domains")
	listCmd.Flags().Int("pages", 10, "How many pages of data to get when using the --keyword option")
	viper.BindPFlags(listCmd.Flags())
}
