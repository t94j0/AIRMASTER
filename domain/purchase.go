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
	return &Domain{url, categorization, 0, 0}
}

func (d *Domain) PromptPurchase() {
	fmt.Println("This feature has not been implemented!")
	return

	clientList := []Registrars{
		godaddy.NewClient(
			viper.GetString("godaddy.key"),
			viper.GetString("godaddy.secret"),
			godaddy.Contact{
				viper.GetString("user.first"),
				viper.GetString("user.middle"),
				viper.GetString("user.last"),
				viper.GetString("user.organization"),
				viper.GetString("user.title"),
				viper.GetString("user.email"),
				viper.GetString("user.phone"),
				viper.GetString("user.fax"),
				viper.GetString("user.mailing"),
			},
		),
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
				d.URL, d.Categorization, price, client.GetName,
			)
			reader := bufio.NewReader(os.Stdin)
			input, _ := reader.ReadString('\n')
			if strings.Contains(input, "Y") || strings.Contains(input, "y") {
				client.Purchase(d.URL)
			}
		}
	}
}
