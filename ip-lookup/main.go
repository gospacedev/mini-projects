package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

type IPInfo struct {
	Query         string  `json:"query"`
	Status        string  `json:"status"`
	Continent     string  `json:"continent"`
	ContinentCode string  `json:"continentCode"`
	Country       string  `json:"country"`
	CountryCode   string  `json:"countryCode"`
	Region        string  `json:"region"`
	RegionName    string  `json:"regionName"`
	City          string  `json:"city"`
	District      string  `json:"district"`
	Zip           string  `json:"zip"`
	Lat           float64 `json:"lat"`
	Lon           float64 `json:"lon"`
	Timezone      string  `json:"timezone"`
	Offset        int     `json:"offset"`
	Currency      string  `json:"currency"`
	Isp           string  `json:"isp"`
	Org           string  `json:"org"`
	As            string  `json:"as"`
	Asname        string  `json:"asname"`
	Mobile        bool    `json:"mobile"`
	Proxy         bool    `json:"proxy"`
	Hosting       bool    `json:"hosting"`
}

type IP struct {
    Query string
}

func PublicIPAddress() string {
    req, err := http.Get("http://ip-api.com/json/")
    if err != nil {
        return err.Error()
    }
    defer req.Body.Close()

    body, err := ioutil.ReadAll(req.Body)
    if err != nil {
        return err.Error()
    }

    var ip IP
    json.Unmarshal(body, &ip)

    return ip.Query
}

func main() {
	localIP := PublicIPAddress()

	var userIP string

	fmt.Println("IP Lookup")
	fmt.Println("Enter an IP address (leave empty for local address):")
	fmt.Scanln(&userIP)

	if userIP == "" {
		userIP += localIP
		fmt.Println(userIP)
	}

	url := "http://ip-api.com/json/" + userIP

	resp, err := http.Get(url)
	if err != nil {
		log.Fatal(err)
	}

	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	var result IPInfo
	if err := json.Unmarshal(body, &result); err != nil {
        fmt.Println("Can not unmarshal JSON")
    }

	fmt.Println("Status: " + result.Status)
	fmt.Println("Continent: " + result.Continent)
	fmt.Println("Country: " + result.Country)
	fmt.Println("Region: " + result.RegionName)
	fmt.Println("City: " + result.City)
	fmt.Println("Timezone: " + result.Timezone)
}
