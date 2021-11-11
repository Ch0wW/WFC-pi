package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"net"
	"net/http"
	"os"
	"runtime"
	"time"
)

var (
	buffer []string
)

type DNSlist struct {
	Domain string `json:"address"`
	Newip  string `json:"newip"`
}

type DNSConfig struct {
	Version  int       `json:"version"`
	DNS      []string  `json:"dns"`
	Redirect []DNSlist `json:"redirections"`
}

func isIPValid(ip string) bool {
	if net.ParseIP(ip) != nil {
		return true
	} else {
		return false
	}
}

func main() {

	//
	fmt.Println("============================")
	fmt.Println("==== WFC-pi DNS Updater ====")
	fmt.Println("============================")
	fmt.Println("")

	var filename string
	var dnsfilepath string

	if runtime.GOOS == "windows" {
		flag.StringVar(&filename, "output", "./dnsmasq.alias.conf", "Output of the file.")
	} else {
		flag.StringVar(&filename, "output", "/etc/dnsmasq.alias.conf", "Output of the file.")
	}
	flag.StringVar(&dnsfilepath, "url", "https://raw.githubusercontent.com/Ch0wW/WFC-pi/main/dns.json", "Path to .json updates.")
	flag.Parse()

	f, err := os.OpenFile(filename, os.O_WRONLY|os.O_CREATE|os.O_TRUNC, 0600)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	fmt.Println("> Getting the list of DNS...")

	var myClient = &http.Client{Timeout: 10 * time.Second}

	var cfg = &DNSConfig{}

	r, err := myClient.Get(dnsfilepath)
	if err != nil {
		panic("Error: Unable to access the DNS list file... Exiting.")
	}
	defer r.Body.Close()

	errno := json.NewDecoder(r.Body).Decode(cfg)
	if errno != nil {
		panic("Error decoding the list. There's probably an error coming from the file itself... Exiting.")
	}

	buffer = append(buffer, "# WIIMMFI-DNS\n")
	buffer = append(buffer, "# This file is auto-generated and updated on every reboot.\n")
	buffer = append(buffer, "# ------------------\n\n")

	for i := 0; i < len(cfg.DNS); i++ {
		ip := cfg.DNS[i]

		if isIPValid(ip) {
			buffer = append(buffer, fmt.Sprintf("server=%s\n", ip))
		} else {
			panic(fmt.Sprintf("%s seems like an invalid IP. That's an error from the file, and not from you... Exiting", ip))
		}
		fmt.Printf("server=%s\n", ip)
	}

	buffer = append(buffer, "\n")

	for i := 0; i < len(cfg.Redirect); i++ {
		info := cfg.Redirect[i]

		if isIPValid(info.Newip) {
			buffer = append(buffer, fmt.Sprintf("address=/%s/%s\n", info.Domain, info.Newip))
		} else {
			panic(fmt.Sprintf("Error parsing %s. That's an error from the file, and not from you... Exiting", info.Newip))
		}
	}

	fmt.Println("> Writing configuration...")
	for i := 0; i < len(buffer); i++ {
		if _, err = f.WriteString(buffer[i]); err != nil {
			panic(err)
		}
	}

	fmt.Println("> Update done!")
}
