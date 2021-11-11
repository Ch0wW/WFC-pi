#!/bin/bash

# Modify this to your own country.
ISO_WIFICOUNTRY="FR"

# If you know what you are doing, you can also modify these values!
AP_NETWORK="192.168.50"
DHCP_DURATION="24h"

echo "> Enabling WiFi code according to the region (default: FR)"
sudo raspi-config nonint do_wifi_country $ISO_WIFICOUNTRY

# Enable services
echo "> Enabling hostapd service..."
sudo systemctl unmask hostapd
sudo systemctl enable hostapd

echo "> Adding our Wi-Fi interface to /etc/dhcpcd.conf..."
sudo cat << EOF >> /etc/dhcpcd.conf
interface wlan0
    static ip_address=$AP_NETWORK.1/24
    nohook wpa_supplicant
EOF

echo "> Enabling IPv4 routing for WFC-pi..."
sudo cat > /etc/sysctl.d/routed-ap.conf << EOF
# Enable IPv4 routing
net.ipv4.ip_forward=1
EOF

# Configure NAT router
echo "> Creating NAT routing..."
sudo iptables -t nat -A POSTROUTING -o eth0 -j MASQUERADE
sudo netfilter-persistent save

# Configure DHCP
echo "> Configuring our DHCP server for our access point..."
sudo cat << EOF > /etc/dnsmasq.conf
interface=wlan0 # Listening interface
domain=wlan

#DHCP Settings
dhcp-range=$AP_NETWORK.2,$AP_NETWORK.250,255.255.255.0,$DHCP_DURATION

# Hardcoded alias for wfcpi
address=/wfcpi/$AP_NETWORK.1
no-resolv

# Configuration file for WFC redirections
conf-file=/etc/dnsmasq.alias.conf
EOF

# Configure WiFi hotspot
echo "> Configuring the wfcpi Access Point..."
sudo cat << EOF > /etc/hostapd/hostapd.conf
country_code=$ISO_WIFICOUNTRY
interface=wlan0
ssid=wfcpi
hw_mode=b
channel=7
macaddr_acl=0
ignore_broadcast_ssid=1
EOF

echo ""
echo ""
echo "======================"
echo ""
echo "At this point, it is really important to reboot your Raspberry Pi."
echo ""
echo "If you don't know how to do, type this command >>>> sudo reboot"