package domain

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"os/exec"
	"strconv"
	"strings"
	"time"

	"github.com/spf13/viper"
)

// DomainCategorization is a struct that is returned by bluecoat when asked
// to classify a domain
type DomainCategorization struct {
	// URL of domain
	URL string `json:"url"`
	// Error is returned if there is an error checking the domain
	Error     string `json:"error"`
	ErrorType string `json:"errorType"`
	// Unrated checks if the domain is uncategorized
	Unrated    bool `json:"unrated"`
	TrackingID int  `json:"curtrackingid"`
	Locked     bool `json:"locked"`
	Multiple   bool `json:"multiple"`
	// RateDate is when the domain was last rated
	RateDate string `json:"ratedate"`
	// Categorization gives the categorization of the domain in an `a` tag
	Categorization    string `json:"categorization"`
	ThreatRiskLevel   string `json:"threatrisklevel"`
	ThreatRiskLevelEn string `json:"threatrisklevel_en"`
	Linkable          bool   `json:"linkable"`
}

// Cooldown is how long to wait when "intrusion" is returned
const Cooldown = 5.5

// CheckDomain checks the categorization of a domain and returns a propmt when
// a domain is found that the user might want.
func CheckDomain(domain string, client *http.Client) error {
	// Make a request to query for the specified domain
	cat, err := makeRequest(domain, "", client)
	if err != nil {
		return err
	}

	// Check the error type
	switch cat.ErrorType {
	case "captcha":
		solveCaptcha(domain, client)
		return CheckDomain(domain, client)
	case "intrusion":

		fmt.Fprintf(os.Stderr, "Waiting %d minuites to cool down\n", Cooldown)
		time.Sleep(time.Minute * time.Duration(Cooldown))
		return CheckDomain(domain, client)
	case "":
		// Don't use Unrated domains
		if !cat.Unrated {
			categorization := getCategorization(cat.Categorization)
			// Purchase domain if that option is specified
			if viper.GetBool("purchase") {
				domainURL, err := url.Parse(cat.URL)
				if err != nil {
					return err
				}
				newDomain := NewDomain(domainURL.Host, categorization)
				newDomain.PromptPurchase()
				return nil
			} else {
				fmt.Println("Found:", cat.URL, "-", categorization)
				return nil
			}
		}
	default:
		return errors.New(cat.Error)
	}
	return nil
}

// makeRequest makes a bluecoat domain categorization request and returns a
// DomainCategorization object
func makeRequest(domain, captcha string, client *http.Client) (*DomainCategorization, error) {
	// Set captcha if a captcha is specified
	v := url.Values{}
	if captcha != "" {
		v.Set("captcha", captcha)
	}
	v.Set("url", domain)
	body := bytes.NewBufferString(v.Encode())

	// Create Request object with body specified by `v`
	request, err := http.NewRequest("POST", "https://sitereview.bluecoat.com/rest/categorization", body)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	// Set Headers
	request.Header.Add("User-Agent", "AIRMASTER")
	request.Header.Add("X-Requested-With", "XMLHttpRequest")
	request.Header.Add("Referer", "https://sitereview.bluecoat.com/sitereview.jsp")

	// Make HTTP POST request
	response, err := client.Do(request)
	if err != nil {
		fmt.Println(err)
		return nil, err
	}

	cat := &DomainCategorization{}
	if err := json.NewDecoder(response.Body).Decode(cat); err != nil {
		fmt.Println(err)
		return nil, err
	}

	return cat, nil
}

// solveCaptcha is what makes this project nice. When called, it solved the
// captcha for the IP specified
func solveCaptcha(domain string, client *http.Client) {
	// Get the captcha for the specified time. This needs to happen to get
	// the session cookie
	now := time.Now().Unix()
	url := "https://sitereview.bluecoat.com/rest/captcha.jpg?" + strconv.FormatInt(now, 10)

	// Make HTTP request with specified headers
	request, err := http.NewRequest("GET", url, nil)
	if err != nil {
		fmt.Println("Error creating request:", err)
		return
	}

	request.Header.Add("User-Agent", "AIRMASTER")
	request.Header.Add("X-Requested-With", "XMLHttpRequest")
	request.Header.Add("Referer", "https://sitereview.bluecoat.com/sitereview.jsp")

	response, err := client.Do(request)
	if err != nil {
		fmt.Println("Error making response:", err)
	}

	// Use Tesseract to do OCR from standard input
	cmd := exec.Command("tesseract", "stdin", "stdout")
	cmd.Stdin = response.Body

	// Pipe the standard output to the `out` buffer
	var out bytes.Buffer
	cmd.Stdout = &out

	if err := cmd.Run(); err != nil {
		fmt.Println(err)
		return
	}
	output := strings.TrimSpace(strings.Trim(out.String(), " "))

	// Make request to stop IP blacklist
	dmn, err := makeRequest(domain, output, client)
	if err != nil {
		fmt.Println("Error making request:", err)
		return
	} else if dmn.ErrorType == "captcha" {
		solveCaptcha(domain, client)
	}

	return
}
