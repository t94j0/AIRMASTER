package domain

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/cookiejar"
	"os"
	"strings"
)

// ParseFile takes a path and runs every domain through the CheckDomain function
func ParseFile(filePath string) error {
	file, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	jar, err := cookiejar.New(&cookiejar.Options{})
	if err != nil {
		return err
	}

	client := &http.Client{
		Jar: jar,
	}

	for _, url := range strings.Split(string(file), "\n") {
		if err := CheckDomain(url, client); err != nil {
			fmt.Fprintln(os.Stderr, err)
		}
	}

	return nil
}
