package main

import (
	"encoding/csv"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"os"
)

func main() {
	// request amazon ip ranges data
	res, _ := http.Get("https://ip-ranges.amazonaws.com/ip-ranges.json")
	defer res.Body.Close()
	body, _ := ioutil.ReadAll(res.Body)

	// parse json information
	var aws struct {
		SyncToken  string `json:"syncToken"`
		CreateDate string `json:"createDate"`
		Prefixes   []struct {
			IPPrefix string `json:"ip_prefix"`
			Region   string `json:"region"`
			Service  string `json:"service"`
		} `json:"prefixes"`
	}
	json.Unmarshal(body, &aws)

	// ouput as csv
	writer := csv.NewWriter(os.Stdout)
	defer writer.Flush()
	writer.Write([]string{"service", "region", "network"})
	for _, p := range aws.Prefixes {
		writer.Write([]string{p.Service, p.Region, p.IPPrefix})
	}
}
