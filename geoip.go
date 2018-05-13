package GeoIP

import (
	"encoding/json"
	"net/http"
)

const (
	GEOIP_URL = "https://json.geoiplookup.io/%s"
)

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

func Lookup(string ip) (*LookupResponse, error) {
	client := http.Client{
		Timeout: 10 * time.Second,
	}

	res, err := client.Get(fmt.Sprintf(GEOIP_URL, ip))

	if err != nil {
		// Assume EU in the face of failure
		return nil, err
	}

	defer res.Body.Close()

	lookupResponse := &LookupResponse{}

	err := json.NewDecoder(res.Body).Decode(lookupResponse)

	if err != nil {
		// Assume EU in the face of failure
		return nil, err
	}

	return lookupResponse, nil
}
