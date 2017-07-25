package domain

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	namecheap "github.com/billputer/go-namecheap"
	"github.com/spf13/viper"
)

type ExpiredDomain struct {
	Site          string
	Registrars    string
	Backlinks     int
	PopBacklinks  int
	Birth         int
	ArchiveOrg    int
	SimilarWeb    string
	Dmoz          string
	DNSCom        string
	DNSNet        string
	DNSOrg        string
	DNSDe         string
	TLDRegistered int
	Related       string
	List          string
	Status        string
}

func ParseKeywords(keywords []string, ncClient *namecheap.Client) error {
	pages := viper.GetInt("pages")
	keyword := strings.Join(keywords, " ")

	rawDomains := make([]ExpiredDomain, 0)

	fmt.Println("Getting domains...")

	// Create rawDomains list which comes directly from expireddomains.com
	for i := 0; i <= pages*25; i += 25 {
		tmpDomains, err := makeEDQuery(keyword, i)
		if err != nil {
			fmt.Println("Error parsing ExpiredDomains query:", err)
		}
		rawDomains = append(rawDomains, tmpDomains...)
	}

	fmt.Println("Done getting domains")

	client := http.DefaultClient

	// Only get domains that are not `Bid` or `In Auction`
	for _, domain := range rawDomains {
		if domain.Site != "" &&
			!strings.Contains(domain.Status, "Bid") &&
			!strings.Contains(domain.Status, "Auction") {

			if err := CheckDomain(domain.Site, client, ncClient, 0); err != nil {
				fmt.Println("Error checking domain:", err)
			}
		}
	}

	return nil
}

func makeEDQuery(query string, page int) ([]ExpiredDomain, error) {
	expiredDomainsURL := "https://www.expireddomains.net/domain-name-search/"
	client := &http.Client{}

	queryURL := expiredDomainsURL + "?start=" + strconv.Itoa(page) + "&q=" + query

	request, err := http.NewRequest("GET", queryURL, nil)
	request.Header.Add("Referer", "https://www.expireddomains.net/domain-name-search/")

	response, err := client.Do(request)
	if err != nil {
		return nil, err
	}

	doc, err := goquery.NewDocumentFromReader(response.Body)
	if err != nil {
		return nil, err
	}

	domains := make([]ExpiredDomain, 0)

	doc.Find("table").Each(func(i int, table *goquery.Selection) {

		table.Find("tr").Each(func(i int, tr *goquery.Selection) {
			var domain ExpiredDomain
			tr.Find("td").Each(func(i int, td *goquery.Selection) {
				switch i {
				case 0:
					domain.Site = td.Find("a").First().Text()
				case 1:
					domain.Backlinks, err = strconv.Atoi(td.Find("a").First().Text())
				case 2:
					domain.PopBacklinks, _ = strconv.Atoi(td.Text())
				case 3:
					domain.Birth, _ = strconv.Atoi(td.Text())
				case 4:
					domain.ArchiveOrg, _ = strconv.Atoi(td.Text())
				case 5:
					domain.SimilarWeb = td.Text()
				case 7:
					domain.Dmoz = td.Text()
				case 8:
					domain.DNSCom = td.Text()
				case 9:
					domain.DNSNet = td.Text()
				case 10:
					domain.DNSOrg = td.Text()
				case 11:
					domain.DNSDe = td.Text()
				case 12:
					domain.TLDRegistered, _ = strconv.Atoi(td.Text())
				case 13:
					domain.Related = td.Text()
				case 14:
					domain.List = td.Text()
				case 15:
					domain.Status = td.Text()
				}
			})
			domains = append(domains, domain)
		})
	})

	return domains, nil
}
