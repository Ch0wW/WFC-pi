#!/bin/bash

echo "> Updating and upgrading packages"
apt update && apt upgrade -y

echo "> Installing hostapd..."
apt install -y hostapd 

echo "> Installing DNSmasq..."
apt install -y dnsmasq

echo "> Installing IPTables-Persistent & netfilter-persistent..."
DEBIAN_FRONTEND=noninteractive apt install -y netfilter-persistent iptables-persistent

echo ""
echo ""
echo "======================"
echo "Packages required are installed!"
echo ""
echo "At this point, it is really important to reboot your Raspberry Pi in order to continue."
echo ""
echo "If you don't know how to do, type this command >>>> sudo reboot"