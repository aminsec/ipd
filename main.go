package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"strconv"
)

type ApiResponse struct {
	Status        string `json:"status"`
	StatusMessage string `json:"status_message"`
	Data          struct {
		IP        string `json:"ip"`
		PTRRecord string `json:"ptr_record"`
		Prefixes  []struct {
			Prefix string `json:"prefix"`
			IP     string `json:"ip"`
			CIDR   int    `json:"cidr"`
			ASN    struct {
				ASN         int    `json:"asn"`
				Name        string `json:"name"`
				Description string `json:"description"`
				CountryCode string `json:"country_code"`
			} `json:"asn"`
			Name        interface{} `json:"name"`
			Description string      `json:"description"`
			CountryCode string      `json:"country_code"`
		} `json:"prefixes"`
		RIRAllocation struct {
			RIRName          string      `json:"rir_name"`
			CountryCode      interface{} `json:"country_code"`
			IP               string      `json:"ip"`
			CIDR             int         `json:"cidr"`
			Prefix           string      `json:"prefix"`
			DateAllocated    string      `json:"date_allocated"`
			AllocationStatus string      `json:"allocation_status"`
		} `json:"rir_allocation"`
		IANAAssignment struct {
			AssignmentStatus string      `json:"assignment_status"`
			Description      string      `json:"description"`
			WhoisServer      string      `json:"whois_server"`
			DateAssigned     interface{} `json:"date_assigned"`
		} `json:"iana_assignment"`
		MaxMind struct {
			CountryCode string      `json:"country_code"`
			City        interface{} `json:"city"`
		} `json:"maxmind"`
	} `json:"data"`
	Meta struct {
		TimeZone      string `json:"time_zone"`
		ApiVersion    int    `json:"api_version"`
		ExecutionTime string `json:"execution_time"`
	} `json:"@meta"`
}

func main() {

	var ip string
	fmt.Scan(&ip)
	url := "https://api.bgpview.io/ip/"
	url += ip

	resp, _ := http.Get(url)
	body, _ := ioutil.ReadAll(resp.Body)
	data := string(body)

	var informaiton ApiResponse

	err := json.Unmarshal([]byte(data), &informaiton)
	if err != nil {
		fmt.Printf("could not unmarshal json: %s\n", err)
		return
	}

	output := map[string]string{}
	output["IP"] = informaiton.Data.IP

	ASN_Info := map[string]string{}
	for _, info := range informaiton.Data.Prefixes {
		ASN_Info["Num"] = strconv.Itoa(info.ASN.ASN)
		ASN_Info["Name"] = info.ASN.Name
		ASN_Info["Desc"] = info.ASN.Description
		output["CIDR"] = ip + "/" + strconv.Itoa(info.CIDR)
		output["CDesc"] = info.Description

	}

	fmt.Println("IP" + ": " + output["IP"])
	fmt.Println("ASN: ")
	for key, vals := range ASN_Info {
		fmt.Printf("   " + key + ": " + vals + "\n")
	}

	fmt.Println("CIDR: ")
	fmt.Printf("   " + "Range: " + output["CIDR"] + "\n")
	fmt.Printf("   " + "Desc: " + output["CDesc"] + "\n")

}
