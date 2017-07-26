package domain

import (
	"bufio"
	"errors"
	"fmt"
	"os"
	"strings"

	namecheap "github.com/billputer/go-namecheap"
	"github.com/spf13/viper"
	"github.com/t94j0/godaddy"
)

type Domain struct {
	URL            string
	Categorization string
	GodaddyPrice   uint64
	NamecheapPrice uint64
}

var ErrUnavailable = errors.New("Domain unavailable")

func NewDomain(url, categorization string) *Domain {
	return &Domain{url, categorization, 0, 0}
}

func (d *Domain) PromptPurchase() {
	fmt.Println("This feature has not been implemented!")
	return

	var godaddyClient *godaddy.Client

	if viper.GetBool("usingGodaddy") {
		godaddyClient = godaddy.NewClient(viper.GetString("godaddy.key"), viper.GetString("godaddy.secret"))
	}

	if d.isAvailable(godaddyClient) {
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

func (d *Domain) isAvailable(g *godaddy.Client) bool {
	available, price, _ := g.IsAvailable(d.URL)
	if available {
		d.GodaddyPrice = price
	}
	return available
}

func (d *Domain) purchase(client *namecheap.Client) {
	if _, err := client.DomainsGetList(); err != nil {
		fmt.Println(err)
	}
}
