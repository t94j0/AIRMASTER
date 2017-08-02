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
4. Copy the configuration found in the Config section and place it in `~/.AIRMASTER.json`


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
	* phone - Phone number (format: +[country_code].XXXXXXXXXX. Ex: +1.9999999999)
	* fax - Fax number
	* address
	* city
	* postal
	* country_code - ISO ["Alpha 2 Code"](http://www.nationsonline.org/oneworld/country_code_list.htm)
* godaddy - Godaddy configuration
	* godaddyKey
	* godaddySecret
* namecheap - Namecheap configuration (Not built yet!)
	* namecheapUser 
	* namecheapKey
	* namecheapUsername
* file - Sets location for file to check domains from
* keyword - Set keywords

(*) is required

### Example Config

Before anyone freaks out, the API key is a test key taken from [the GoDaddy docs](https://developer.godaddy.com/doc)

```
{
    "godaddyKey": "UzQxLikm_46KxDFnbjN7cQjmw6wocia",
    "godaddySecret": "46L26ydpkwMaKZV6uVdDWe",
    "first": "Max",
    "last": "Harley",
    "organization": "Max Co.",
    "title": "CEO",
    "email": "maxh@maxh.io",
    "phone": "+1.9999999",
    "address": "1 Awesome Dr.",
    "city": "Charleston",
    "state": "SC",
    "postal": "2946X",
    "country_code": "US"
}
```
