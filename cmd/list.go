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
		// Check namecheap and godaddy keys
		godaddyKey, godaddySecret := viper.GetString("godaddy.key"), viper.GetString("godaddy.secret")
		if godaddyKey == "" || godaddySecret == "" {
			fmt.Println("Not using GoDaddy")
		} else {
			fmt.Println("Using GoDaddy")
			viper.Set("usingGodaddy", true)
		}

		ncUser, ncKey, ncUsername := viper.GetString("namecheap.user"), viper.GetString("namecheap.key"), viper.GetString("namecheap.username")
		if ncUser == "" || ncKey == "" || ncUsername == "" {
			fmt.Println("Not using Namecheap")
		} else {
			fmt.Println("Using Namecheap")
			viper.Set("usingNamecheap", true)
		}

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

	viper.SetDefault("usingGodaddy", false)
	viper.SetDefault("usingNamecheap", false)
}
