package cmd

import (
	"fmt"
	"os"

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
		checkConfig()
		switch {
		// Configure domain finding mechanism
		case viper.GetString("file") != "":
			if err := domain.ParseFile(viper.GetString("file")); err != nil {
				fmt.Println(err)
			}
		case len(viper.GetStringSlice("keyword")) != 0:
			if err := domain.ParseKeywords(viper.GetStringSlice("keyword")); err != nil {
				fmt.Println(err)
			}
		default:
			fmt.Println("Please specify either a file or keywords")
			fmt.Println(cmd.Usage())
		}
	},
}

func checkConfig() {
	godaddyConfig := []string{
		"godaddyKey",
		"godaddySecret",
		"first",
		"last",
		"organization",
		"title",
		"email",
		"phone",
		"address",
		"city",
		"state",
		"postal",
		"country_code",
	}

	for _, c := range godaddyConfig {
		if viper.GetString(c) == "" {
			fmt.Printf("Not using GoDaddy (Provide %s)\n", c)
			break
		}
	}
}

func init() {
	RootCmd.AddCommand(listCmd)

	listCmd.Flags().BoolP("purchase", "p", false, "Purchase domains that are listed")
	listCmd.Flags().Bool("expiredcom", false, "Uses the current list of deleted .com domains from expiredomains.com")
	listCmd.Flags().StringP("file", "f", "", "File used for checking domains")
	listCmd.Flags().StringSliceP("keyword", "k", nil, "Keyword for searching domains")
	listCmd.Flags().Int("pages", 10, "How many pages of data to get when using the --keyword option")
	if err := viper.BindPFlags(listCmd.Flags()); err != nil {
		fmt.Fprintln(os.Stderr, "Error binding flags:", err)
		os.Exit(1)
	}

}
