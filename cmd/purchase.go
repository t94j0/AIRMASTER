package cmd

import (
	"bufio"
	"fmt"
	"net/url"
	"os"
	"strings"

	"github.com/spf13/cobra"
	"github.com/spf13/viper"
	"github.com/t94j0/AIRMASTER/domain"
)

// purchaseCmd represents the purchase command
var purchaseCmd = &cobra.Command{
	Use:   "purchase",
	Short: "Purchase domains straight after listing",
	Long: `Purchase should be used if you just want to purchase one domain or from a
previous listing. This can be more efficent if you want to leave the "list"
command running for a long time.

Example: AIRMASTER purchase --domain "maxh.io"
         or
         AIRMASTER list --file ./domains.txt > categorized_domains.txt
	 cat categorized_domains.txt | AIRMASTER purchase`,
	Run: func(cmd *cobra.Command, args []string) {
		if viper.GetString("domain") != "" {
			singleDomain := domain.NewDomain(viper.GetString("domain"), "")
			singleDomain.PromptPurchase()
		} else if viper.GetString("list") != "" {
			file, err := os.Open(viper.GetString("list"))
			if err != nil {
				fmt.Fprintln(os.Stderr, "Could not open file:", err)
				return
			}
			scanner := bufio.NewScanner(file)

			for scanner.Scan() {
				// Remove final newline character
				urlLine := strings.Trim(scanner.Text(), "\n")
				// Get just the URL (http://maxh.io)
				urlCat := strings.Split(urlLine, "-")
				if len(urlCat) != 2 {
					fmt.Printf("URL \"%s\" is malformed. Please use output from `AIRMASTER list`\n", urlLine)
					continue
				}
				urlLine = strings.TrimRight(urlCat[0], " ")
				categorization := strings.TrimLeft(urlCat[1], " ")
				urlObj, err := url.Parse(urlLine)
				if err != nil {
					fmt.Println("Failed parsing URL:", urlLine)
					continue
				}
				singleDomain := domain.NewDomain(urlObj.Host, categorization)
				singleDomain.PromptPurchase()
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(purchaseCmd)
	purchaseCmd.Flags().StringP("domain", "d", "", "Used to purchase a single domain")
	purchaseCmd.Flags().StringP("list", "l", "", "Specify output from `list` to get a purchase prompt")

	if err := viper.BindPFlags(purchaseCmd.Flags()); err != nil {
		fmt.Println("Error binding flags:", err)
		os.Exit(1)
	}
}
