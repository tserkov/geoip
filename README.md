# GeoIP
A Golang wrapper for [geoiplookup.io](https://geoiplookup.io/)'s free API.

## Methods
```golang
Lookup(string ip) (*LookupResult, error)
```

## Response
```golang
type LookupResponse struct {
	ASN            string `json:"asn"`
	CountryName    string `json:"country_name"`
	ContinentCode  string `json:"continent_code"`
	CurrencyCode   string `json:"currency_code"`
	CurrencyName   string `json:"currency_name"`
	City           string `json:"city"`
	CountryCode    string `json:"country_code"`
	ConnectionType string `json:"connection_type"`
	District       string `json:"district"`
	Hostname       string `json:"hostname"`
	ISP            string `json:"isp"`
	IP             string `json:"ip"`
	Latitude       string `json:"latitude"`
	Longitude      string `json:"longitude"`
	Org            string `json:"org"`
	PostalCode     string `json:"postal_code"`
	Region         string `json:"region"`
	Success        bool   `json:"success"`
	TimezoneName   string `json:"timezone_name"`
}
```
