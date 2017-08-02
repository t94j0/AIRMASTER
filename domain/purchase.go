package domain

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/spf13/viper"
	"github.com/t94j0/godaddy"
)

// Registrars is an interface for purchasing domain names
type Registrars interface {
	// GetName returns the name of the registrar
	GetName() string

	// IsAvailable returns if the domain is available and the price
	IsAvailable(domain string) (bool, uint64, error)

	// Purchase takes a domains and buys it
	Purchase(domain string) error
}

// Domain is a description of the domain with categorization
type Domain struct {
	URL            string
	Categorization string
}

// NewDomain creates a new Domain struct that can be used to check availability
// and purchase domains
func NewDomain(url, categorization string) *Domain {
	return &Domain{url, categorization}
}

// PromptPurchase is a CUI for purchasing a domain. It uses the helpers given
// to actually purchase it.
func (d *Domain) PromptPurchase() {
	// Get all registrars that user has enabled
	allRegistrars := getRegistrars()

	// Get all registrars that have the domain available
	availableRegistrars, prices := getAvailability(d.URL, allRegistrars)
	if len(availableRegistrars) == 0 {
		return
	}

	// Give the user options for how to purchase the domain, or the option not to
	fmt.Println("-1. Do not purchase")

	for i, registrar := range availableRegistrars {
		fmt.Printf(
			"%d. %s (%s) available on %s for $%d\n",
			i,
			d.URL,
			d.Categorization,
			registrar.GetName(),
			prices[i],
		)
	}

	// UI for purchasing the domain
	for {
		fmt.Printf("Choose an option: ")
		reader := bufio.NewReader(os.Stdin)
		input, err := reader.ReadString('\n')
		if err != nil {
			fmt.Fprintln(os.Stderr, err)
			break
		}
		input = strings.Trim(input, "\n")
		choice, err := strconv.Atoi(input)
		if err != nil {
			fmt.Println("Please input a number...")
			continue
		}
		if len(availableRegistrars) < choice || choice < -1 {
			fmt.Println("Not a choice")
			continue
		}

		if choice == -1 {
			break
		}

		if err := availableRegistrars[choice].Purchase(d.URL); err != nil {
			fmt.Fprintln(os.Stderr, "Error purchasing domain:", err)
			break
		}

		fmt.Println("Success!")
		break
	}
	fmt.Printf("\n\n")
}
func getRegistrars() []Registrars {
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
	}

	return clientList
}

func getAvailability(domain string, allRegistrars []Registrars) ([]Registrars, []uint64) {
	registrars := make([]Registrars, 0)
	prices := make([]uint64, 0)

	for _, registrar := range allRegistrars {
		isAvailable, price, err := registrar.IsAvailable(domain)
		if err != nil {
			fmt.Fprintln(os.Stderr, "Error checking availablity:", err)
			continue
		}
		if isAvailable {
			registrars = append(registrars, registrar)
			prices = append(prices, price)
		}
	}

	return registrars, prices
}
