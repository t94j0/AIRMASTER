# AIRMASTER
(Thanks [NSA Name-O-Matic](https://divergentdave.github.io/nsa-o-matic/) for the name)

## How this project is different than [domain hunter](https://github.com/minisllc/domainhunter)

* Actively maintained
* Bypasses captcha for [bluecoat.com](https://sitereview.bluecoat.com/sitereview.jsp) by using OCR because their captcha is shit
* Plans for purchasing domain with Namecheap and GoDaddy

## Installation

### macOS

1. Install [homebrew](`https://brew.sh/`)
2. In your terminal, run `brew install tesseract`
3. Build Go project or download a release


## How to use

Right now, this only supports listing domains that a red team might want to purchase. Although, you can do it one of two ways:

1. With a domain list file

`AIRMASTER list --file ./path/to/file.txt`

2. With keywords

`AIRMASTER list --keyword max --keyword cool`

If multiple keywords are specified, they are combined by AND, so in the example above, you will get `maxiscool.com, max-is-kinda-cool.com, cool-memes-to-the-max.com`

The help should be very obvious, so if you are stuck, try using `AIRMASTER --help`

## Config

You can access the configuration by editing the `~/.AIRMASTER.yaml` file.

The options are:
* (*) user - Used for whois data
	* first - Your first name
	* middle - Your middle name
	* last - Your last name
	* organization - Organization that you belong to
	* title - Title at organization
	* email - Email for contact
	* phone - Phone number (format: TODO)
	* fax - Fax number
	* address
	* city
	* postal
	* country_code - ISO ["Alpha 2 Code"](http://www.nationsonline.org/oneworld/country_code_list.htm)
* godaddy - Godaddy configuration
	* key
	* secret
* namecheap - Namecheap configuration
	* user
	* key
	* username
* file - Sets location for file to check domains from
* keyword - Set keywords

(*) is required
