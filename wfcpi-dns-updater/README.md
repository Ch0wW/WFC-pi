# WFCPI-DNS-Updater

* Written in: Golang

This tool automatically checks this repository (and more precisely, `dns.json`) and updates the configuration file on DNSMASQ.

# About Linux/Raspberry Pi
This tool REQUIRES root access to overwrite `/etc/dnsmasq.alias.conf` (a file that is generated specifically for WFC-pi) and properly update the DNS used for the servers.

# Parameters
- `-output` : Output file for the configuration file. 
- `-url` : URL of the json config.

# About Windows
This tool is only used to check if the output file is properly parsed, by writing `dnsmasq.alias.conf` in the folder it's been run.