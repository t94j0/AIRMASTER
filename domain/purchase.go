package domain

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	"github.com/spf13/viper"
	"github.com/t94j0/godaddy"
)

type Registrars interface {
	// GetName returns the name of the registrar
	GetName() string

	// IsAvailable returns if the domain is available and the price
	IsAvailable(domain string) (bool, uint64, error)

	// Purchase takes a domains and buys it
	Purchase(domain string) error
}

type Domain struct {
	URL            string
	Categorization string
}

var ErrUnavailable = errors.New("Domain unavailable")

func NewDomain(url, categorization string) *Domain {
	return &Domain{url, categorization}
}

func (d *Domain) PromptPurchase() {
	var clientList []Registrars

	godaddyClient, err := godaddy.NewClient(
		viper.GetString("godaddyKey"),
		viper.GetString("godaddySecret"),
		godaddy.Contact{
			viper.GetString("first"),
			viper.GetString("middle"),
			viper.GetString("last"),
			viper.GetString("organization"),
			viper.GetString("title"),
			viper.GetString("email"),
			viper.GetString("phone"),
			viper.GetString("fax"),
			godaddy.Address{
				viper.GetString("address"),
				viper.GetString("city"),
				viper.GetString("state"),
				viper.GetString("postal"),
				viper.GetString("country_code"),
			},
		},
	)

	// TODO: Have a more elegant way to handle this
	if err == nil {
		clientList = append(clientList, godaddyClient)
	} else {
		fmt.Println(err)
	}

	for _, client := range clientList {
		isAvailable, price, err := client.IsAvailable(d.URL)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error getting availability:", err)
			continue
		}

		if isAvailable {
			fmt.Printf(
				"Would you like to purchase \"%s\" (%s) for %d from %s (y/N): ",
				d.URL, d.Categorization, price, client.GetName(),
			)
			reader := bufio.NewReader(os.Stdin)
			input, _ := reader.ReadString('\n')
			if strings.Contains(input, "Y") || strings.Contains(input, "y") {
				if err := client.Purchase(d.URL); err != nil {
					fmt.Fprintln(os.Stderr, "Error purchasing domain:", err)
				}
			}
		} else {
			fmt.Printf("Found %s, but not available\n", d.URL)
		}
	}
}
