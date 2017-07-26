package domain

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	namecheap "github.com/billputer/go-namecheap"
)

type Domain struct {
	URL            string
	Categorization string
}

var ErrUnavailable = errors.New("Domain unavailable")

func NewDomain(url, categorization string) *Domain {
	return &Domain{url, categorization}
}

func (d *Domain) PromptPurchase() {
	fmt.Println("This feature has not been implemented!")
	return

	if d.isAvailable() {
		fmt.Println("Domain", d.URL, "("+d.Categorization+")"+" is available")
		fmt.Printf("Would you like to purchase (y/N): ")
		reader := bufio.NewReader(os.Stdin)
		input, _ := reader.ReadString('\n')
		if strings.Contains(input, "Y") || strings.Contains(input, "y") {
			d.purchase(nil)
		} else {
			return
		}
	}
}

func (d *Domain) isAvailable() bool {
	return true
}

func (d *Domain) purchase(client *namecheap.Client) {
	if _, err := client.DomainsGetList(); err != nil {
		fmt.Println(err)
	}
}
